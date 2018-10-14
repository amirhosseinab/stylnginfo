<template>
    <svg id="viewer" ref="svg">

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
                    .range(d3.schemeSpectral[11]);

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

                const leaf = svg.selectAll("g")
                    .data(root.leaves(), d => d.data.fileName)
                    .enter().append("g")
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