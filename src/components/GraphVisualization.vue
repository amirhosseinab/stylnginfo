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
                    textRotation: -30
                },
            }
        },
        computed: {
            ...mapGetters(['graphData', 'htmlFiles', 'cssFiles', 'analyzed']),
            selectedHtmlFiles() {
                return this.htmlFiles.filter(hf => hf.selected).map(hf => hf.name);
            },
            selectedCssFiles() {
                return this.cssFiles.filter(cf => cf.selected).map(cf => cf.name);
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
                let margin = {top: 30, right: 30, bottom: 120, left: 200};

                let xScale = d3.scaleBand()
                        .domain(this.selectedCssFiles)
                        .range([margin.left, this.settings.width - margin.right])
                        .paddingInner(1)
                        .paddingOuter(.8),
                    yScale = d3.scaleBand()
                        .domain(this.selectedHtmlFiles)
                        .range([margin.top, this.settings.height - margin.bottom])
                        .paddingInner(1)
                        .paddingOuter(.8);

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
                    .attr("transform", d => "rotate(" + this.settings.textRotation + " " + (margin.left - 5) + " " + yScale(d) + ")")
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
                    .attr("dy", "0.5em")
                    .attr("dx", "-0.3em")
                    .attr("transform", d => "rotate(" + this.settings.textRotation + " " + xScale(d) + " " + (this.settings.height - margin.bottom) + ")")
                    .attr("class", "text css");
                cssTick.append("line")
                    .attr("class", "line css")
                    .attr("x1", d => xScale(d))
                    .attr("y1", margin.top)
                    .attr("x2", d => xScale(d))
                    .attr("y2", this.settings.height - margin.bottom);

                cssTickGroupData.attr("opacity", null);
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
                    stroke: #dfdfdf;
                    stroke-width: 2;
                }
                text {
                    &.html {
                        fill: $blue-color;
                        font-size: 1em;
                    }
                    &.css {
                        fill: $green-color;
                        font-size: 1.2em;
                    }
                }
            }
        }
    }
</style>

