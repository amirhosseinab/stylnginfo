<template>
    <div class="result-pane">
        <div class="toolbar graph">
            <div class="toolbar-item">

            </div>
        </div>
        <div class="graph-pane">
            <svg id="viewer" :height="settings.height" :width="settings.width">
                <g class="graph"></g>
            </svg>
        </div>
    </div>
</template>

<script>
    import {mapGetters} from 'vuex'
    import * as d3 from 'd3'

    export default {
        data() {
            return {
                settings: {
                    width: 1100,
                    height: 750,
                    removedOpacity: .2,
                    htmlTextRotation: -30,
                    cssTextRotation: -60,
                },
            }
        },
        computed: {
            ...mapGetters(['graphData', 'htmlFiles', 'cssFiles', 'fileLinks', 'selectors', 'htmlToSelectorLinks', 'cssToSelectorLinks', 'analyzed']),
            selectedHtmlFiles() {
                return this.htmlFiles.filter(hf => hf.selected).map(hf => hf.name);
            },
            selectedCssFiles() {
                return this.cssFiles.filter(cf => cf.selected).map(cf => cf.name);
            },
            selectedFileLinks() {
                return this.fileLinks
                    .filter(fl => this.selectedHtmlFiles.includes(fl.source) && this.selectedCssFiles.includes(fl.target))
            },
            htmlSelectors() {
                return this.selectedHtmlFiles.map(h => {
                    return {
                        htmlFileName: h, cssFiles: this.selectedFileLinks
                            .filter(sf => sf.source === h)
                            .map(sf => {
                                return {
                                    name: sf.target, selectorCount: this.cssToSelectorLinks
                                        .filter(cs => cs.source === sf.target)
                                        .filter(cs => this.htmlToSelectorLinks
                                            .find(hs => hs.source === h && hs.target === cs.target))
                                        .length
                                }
                            })
                    }
                });
            }
        },
        watch: {
            selectedHtmlFiles() {
                this.renderGraph()
            },
            selectedCssFiles() {
                this.renderGraph()
            }
        },
        methods: {
            renderGraph() {
                let v = d3.select("#viewer");
                let margin = {top: 30, right: 30, bottom: 180, left: 200};

                let maxRadiusDomain = Math.max(...this.htmlSelectors
                    .map(hs => hs.cssFiles
                        .map(cf => cf.selectorCount)
                        .reduce((a, b) => Math.max(a, b))));


                let xScale = d3.scaleBand()
                        .domain(this.selectedCssFiles)
                        .range([margin.left, this.settings.width - margin.right])
                        .paddingInner(1)
                        .paddingOuter(.8),
                    yScale = d3.scaleBand()
                        .domain(this.selectedHtmlFiles)
                        .range([margin.top, this.settings.height - margin.bottom])
                        .paddingInner(1)
                        .paddingOuter(.8),
                    wScale = d3.scaleBand()
                        .domain(this.selectedCssFiles.length > this.selectedHtmlFiles.length ?
                            this.selectedCssFiles : this.selectedHtmlFiles)
                        .range([margin.left, this.settings.width - margin.right])
                        .paddingInner(.4)
                        .paddingOuter(.2),

                    cScale = d3.scaleLinear()
                        .domain([0, maxRadiusDomain])
                        .range(["#CF9BFF", "#ff0a7c"]),
                    // cScale = d3.scaleQuantize()
                    //     .domain([0, maxRadiusDomain])
                    //     .range(["#ff0a7c","#60d440","#E2B257","#ec7f2a","#ff0a7c"]),
                    rScale = d3.scaleSqrt()
                        .domain([0, maxRadiusDomain])
                        .range([5, wScale.bandwidth()]);

                let graph = v.select("g.graph");

                // HTML
                let htmlTickGroupData = graph.selectAll("g.tick.html")
                    .data(this.selectedHtmlFiles, d => d);
                htmlTickGroupData.exit()
                    .attr("opacity", this.settings.removedOpacity);

                let htmlTick = htmlTickGroupData.enter()
                    .append("g")
                    .attr("class", "tick html");
                htmlTick.append("text")
                    .text(d => d)
                    .attr("text-anchor", "end")
                    .attr("x", margin.left - 5)
                    .attr("y", d => yScale(d))
                    .attr("dy", "0.2em")
                    .attr("transform", d => "rotate(" + this.settings.htmlTextRotation + " " + (margin.left - 5) + " " + yScale(d) + ")")
                    .attr("class", "text html");
                htmlTick.append("line")
                    .attr("class", "line html")
                    .attr("x1", margin.left)
                    .attr("y1", d => yScale(d))
                    .attr("x2", this.settings.width - margin.right)
                    .attr("y2", d => yScale(d));

                htmlTickGroupData.attr("opacity", null);

                // CSS
                let cssTickGroupData = graph.selectAll("g.tick.css")
                    .data(this.selectedCssFiles, d => d);
                cssTickGroupData.exit()
                    .attr("opacity", this.settings.removedOpacity);

                let cssTick = cssTickGroupData.enter()
                    .append("g")
                    .attr("class", "tick css");
                cssTick.append("text")
                    .text(d => d)
                    .attr("text-anchor", "end")
                    .attr("x", d => xScale(d))
                    .attr("y", this.settings.height - margin.bottom)
                    .attr("dy", "0.3em")
                    .attr("dx", "-0.2em")
                    .attr("transform", d => "rotate(" + this.settings.cssTextRotation + " " + xScale(d) + " " + (this.settings.height - margin.bottom) + ")")
                    .attr("class", "text css");
                cssTick.append("line")
                    .attr("class", "line css")
                    .attr("x1", d => xScale(d))
                    .attr("y1", margin.top)
                    .attr("x2", d => xScale(d))
                    .attr("y2", this.settings.height - margin.bottom);

                cssTickGroupData.attr("opacity", null);

                // FILE LINKS
                let fileLinksData = graph.selectAll("circle.file-link").data(this.selectedFileLinks, d => d.source + d.target);
                fileLinksData.exit()
                    .attr("opacity", this.settings.removedOpacity);

                fileLinksData.enter()
                    .append("circle")
                    .attr("class", "file-link")
                    .attr("r", d => {
                        let sc = this.htmlSelectors.find(hs => hs.htmlFileName === d.source)
                            .cssFiles.find(cf => cf.name === d.target).selectorCount;
                        return rScale(sc)
                    })
                    .attr("fill", d => {
                        let sc = this.htmlSelectors.find(hs => hs.htmlFileName === d.source)
                            .cssFiles.find(cf => cf.name === d.target).selectorCount;
                        return cScale(sc)
                    })
                    .attr("cx", d => xScale(d.target))
                    .attr("cy", d => yScale(d.source));

                fileLinksData.attr("opacity", null);

            }
        },
        mounted() {
            this.renderGraph()
        },
    }
</script>
<style lang="scss">
    .result-pane {
        height: 100%;
        text-align: center;
        .graph-pane {
            height: calc(100% - #{$toolbar-height});
            position: relative;
            padding: 20px;
            overflow: auto;
            @extend %scrollbar;
            svg {
                vertical-align: middle;
                line.html, line.css {
                    fill: none;
                    stroke: $graph-gray-color;
                    stroke-width: 2;
                }
                text {
                    &.html {
                        fill: $blue-color;
                        font-size: 0.9em;
                        font-weight: bold;
                    }
                    &.css {
                        fill: $green-color;
                        font-size: 1em;
                        font-weight: bold;
                    }
                }
                circle {
                    &.file-link {
                        //fill: $graph-selector-dot-color;
                        fill-opacity: .7;
                        cursor: default;
                        &:hover {
                            fill-opacity: 1;
                            border: solid 2px $graph-gray-color;
                        }
                    }
                }
            }
        }
    }
</style>

