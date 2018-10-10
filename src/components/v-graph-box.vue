<template>
    <div ref="graph" class="graph-container" @click="toggleGraphSize" :class="{'zoomed':zoomed}">
        <h1 class="graph-title">{{title}}</h1>
        <div class="graph-content">
            <slot class="graph-content">
                <div class="not-available">No Data Available</div>
            </slot>
        </div>
    </div>
</template>

<script>
    export default {
        name: "v-graph-pane",
        data() {
            return {
                zoomed: false,
            }
        },
        props: {
            title: {type: String, required: true}
        },
        methods: {
            toggleGraphSize() {
                if (!this.zoomed) {
                    this.zoomed = true;
                    let screenH = window.innerHeight,
                        screenW = window.innerWidth,
                        graphH = this.$refs.graph.clientHeight * 2,
                        graphW = this.$refs.graph.clientWidth * 2,
                        x = (screenW / 2) - (graphW / 2),
                        y = (screenH / 2) - (graphH / 2);

                    this.$refs.graph.style.width = `${graphW}px`;
                    this.$refs.graph.style.height = `${graphH}px`;
                    this.$refs.graph.style.transform = `translate("${x},${y}")`
                } else {
                    this.$refs.graph.style.width = "";
                    this.$refs.graph.style.height = "";
                    this.$refs.graph.style.transform = "";
                    this.zoomed = false;
                }
            },
        }
    }
</script>

<style scoped lang="scss">
    $base-gray-color: darken($white-color, 10%);
    .graph-container {
        flex: 0 0 initial;
        display: flex;
        flex-flow: column wrap;
        align-items: stretch;
        justify-content: flex-start;

        padding: .6rem;
        border: solid 2px $base-gray-color;
        box-shadow: inset 0 0 30px 5px lighten($base-gray-color, 5%);
        border-radius: 8px;
        background-color: $white-color;

        height: $graph-height;
        width: $graph-width;
        margin: 1rem;
        color: $dark-gray-color;
        z-index: 10;
        cursor: pointer;
        transition: all .15s ease-in;

        &:hover {
            transform: scale(1.01);
            box-shadow: inset 0 0 30px 5px lighten($base-gray-color, 5%), 0 0 8px 2px lighten($base-gray-color, 5%);
        }
        &.zoomed {
            box-shadow: inset 0 0 30px 5px lighten($base-gray-color, 5%);
            position: absolute;
            z-index: 80;
        }
        .graph-title {
            flex: 0 1 auto;
            font-size: 1.3rem;
            font-weight: bold;
            text-align: left;
            margin: 0;
            padding: .2rem .5rem .5rem;

            color: darken($base-gray-color, 10%);
            border-bottom: solid 3px $base-gray-color;
        }
        .graph-content {
            flex: 1 0 auto;
            display: flex;
            flex-flow: row wrap;
            justify-items: stretch;
            align-items: stretch;

            position: relative;
            & > * {
                height: auto;
                width: 100%;
            }
        }

        .not-available {
            flex: 1 0 auto;

            font-size: 2.5rem;
            font-weight: bold;
            color: $base-gray-color;
            align-self: stretch;

            display: flex;
            flex-flow: row wrap;
            align-items: center;
            justify-content: center;
        }
    }
</style>