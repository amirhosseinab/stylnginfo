<template>
    <svg id="viewer" ref="svg">
        <g class="legend"></g>
    </svg>
</template>

<script>
    import * as d3 from "d3";

    export default {
        name: "v-graph-css-files-weight",
        methods: {
            renderGraph() {
                const svg = d3.select("#viewer")
                    .attr("text-anchor", "middle");

                const color = d3.scaleOrdinal()
                    .range(d3.schemeRdPu[5]);

                const height = this.$refs.svg.clientHeight,
                    width = this.$refs.svg.clientWidth;

                const pack = data => d3.pack()
                    .size([width, height])
                    .padding(14)(
                        d3.hierarchy({children: data})
                            .sum(d => d.selectorCount)
                    );
                const fontSize = function (d) {
                    let ebox = this.getBBox(),
                        pbox = this.parentNode.getBBox(),
                        scale = Math.min(pbox.width / ebox.width, pbox.height / ebox.height);
                    d.fontScale = scale;
                };
                const root = pack(this.data);

                const leaf = svg.selectAll("g.file")
                    .data(root.leaves(), d => d.data.fileName)
                    .enter().append("g")
                    .attr("class", "file")
                    .attr("transform", d => `translate(${d.x},${d.y})`);

                leaf.append("circle")
                    .attr("r", d => d.r + 5)
                    .attr("fill-opacity", 0.7)
                    .attr("fill", d => d.data.isUsed ? color(d.data.fileName) : "#fff")
                    .attr("stroke", d => d.data.isUsed ? null : "#7b7b7b")
                    .attr("stroke-width", d => d.data.isUsed ? null : 1);

                leaf.append("text")
                    .text(d => d.data.fileName)
                    .style("font-size", "1px")
                    .each(fontSize)
                    .style("font-size", d => `${d.fontScale * 0.7}px`)
                    .style("fill", "#4a4a4a");

                leaf.append("title")
                    .text(d => `${d.data.fileName}\nSelectors count: ${d.data.selectorCount}`);

                let legend = svg.select("g.legend");
                legend.append("rect")
                    .attr("x", 10)
                    .attr("y", 10)
                    .attr("width", 140)
                    .attr("height", 90)
                    .attr("fill-opacity", 0.7)
                    .attr("fill", "#f9f9f9")
                    .attr("stroke", "#a1a1a1")
                    .attr("stroke-width", 1);
                legend.append("rect")
                    .attr("x", 20)
                    .attr("y", 20)
                    .attr("width", 120)
                    .attr("height", 30)
                    .attr("fill", "#ffffff")
                    .attr("stroke", "#7b7b7b")
                    .attr("stroke-width", .5);
                legend.append("text")
                    .text(`Unused Files: ${this.data.filter(f => !f.isUsed).length}`)
                    .attr("x", 80)
                    .attr("y", 39)
                    .attr("fill", "#7b7b7b")
                    .style("font-size", ".7rem");

                let grad = legend.append("defs")
                    .append("linearGradient");
                grad.attr("id", "legendGrad")
                    .attr("x1", "0%")
                    .attr("y1", "0%")
                    .attr("x2", "100%")
                    .attr("y2", "0%");
                grad.append("stop")
                    .attr("offset", "0%")
                    .style("stop-color", "#fbb4b9")
                    .style("stop-opacity", "1");
                grad.append("stop")
                    .attr("offset", "100%")
                    .style("stop-color", "#7a0177")
                    .style("stop-opacity", "1");

                legend.append("rect")
                    .attr("x", 20)
                    .attr("y", 60)
                    .attr("width", 120)
                    .attr("height", 30)
                    .attr("fill", "url(#legendGrad)")
                    .attr("stroke", "#7b7b7b")
                    .attr("stroke-width", .5);
                legend.append("text")
                    .text(`Used Files: ${this.data.filter(f => f.isUsed).length}`)
                    .attr("x", 80)
                    .attr("y", 79)
                    .attr("fill", "#f1f1f1")
                    .style("font-size", ".7rem")
            }
        },
        props: {
            data: {
                type: Object,
                required: true,
            }
        },
        mounted() {
            this.renderGraph()
        }
    }
</script>

<style scoped lang="scss">

</style>