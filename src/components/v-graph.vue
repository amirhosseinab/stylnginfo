<template>
    <div class="graph-container">

    </div>
</template>

<script>
    export default {
        name: "v-graph-pane"
    }
</script>

<style scoped lang="scss">
    .graph-container {
        flex: 1 1 50vw;

        /*border: solid 1px black;*/

        display: flex;
        align-items: flex-start;
        justify-content: center;

        color: $dark-gray-color;
        z-index: 10;
    }
</style>