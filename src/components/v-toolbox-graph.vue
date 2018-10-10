<template>
    <div class="graph-container">
        <div v-for="graph in graphs" :key="graph.name" @click="graph.show= !graph.show"
             class="graph-selector" :class="{'selected':graph.show}">
            {{graph.title}}
        </div>
    </div>
</template>

<script>
    import {mapGetters} from 'vuex';

    export default {
        name: "v-toolbox-graph",
        computed: {
            ...mapGetters(['graphs'])
        }
    }
</script>

<style scoped lang="scss">
    .graph-container {
        display: flex;
        flex-flow: row wrap;
        align-items: flex-start;
        justify-content: stretch;
        padding: 1rem;
        .graph-selector {
            flex: 1 0 4rem;
            display: flex;
            justify-content: center;
            align-items: center;
            min-height: 3rem;
            min-width: max-content;
            padding: .5rem;
            background-color: $dark-gray-color;
            color: $white-color;
            margin: .5rem;
            border-radius: 8px;
            cursor: pointer;
            &.selected {
                background-color: lighten($mid-gray-color, 10%);
            }
        }
    }
</style>