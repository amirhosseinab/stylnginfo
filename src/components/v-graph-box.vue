<template>
    <div class="graph-container">
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
        props: {
            title: {type: String, required: true}
        },
    }
</script>

<style scoped lang="scss">
    $base-gray-color: darken($white-color, 10%);
    .graph-container {
        position: relative;
        flex: 0 0 initial;
        display: flex;
        flex-flow: column wrap;
        align-items: stretch;
        justify-content: flex-start;

        padding: .6rem;
        border: solid 2px $base-gray-color;
        box-shadow: inset 0 0 30px 5px lighten($base-gray-color, 5%);
        border-radius: 8px;

        height: $graph-height;
        width: $graph-width;
        margin: 1rem 1rem .8rem;
        color: $dark-gray-color;
        z-index: 10;
        cursor: default;
        transition: all .15s ease-in;

        &:hover {
            transform: scale(1.01);
            box-shadow: inset 0 0 30px 5px lighten($base-gray-color, 5%), 0 0 8px 2px lighten($base-gray-color, 5%);
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