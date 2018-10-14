<template>
    <div class="graph-container">
        <h1 class="graph-title">{{title}}</h1>
        <div class="graph-content">
            <slot name="graph">
            </slot>
        </div>
    </div>
</template>

<script>
    export default {
        name: "v-graph-pane",
        data() {
            return {}
        },
        props: {
            title: {type: String, required: true}
        },
        methods: {}
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
        cursor: default;

        .graph-title {
            flex: 0 1 auto;
            font-size: 1rem;
            font-weight: bold;
            text-align: center;
            margin: 0;
            padding: .2rem .5rem .5rem;

            color: darken($base-gray-color, 30%);
            border-bottom: solid 3px darken($base-gray-color, 10%);
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
    }
</style>