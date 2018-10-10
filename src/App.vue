<template>
    <div>
        <main>
            <v-toolbox/>
            <div class="main-content">
                <v-graph-box v-for="graph in activeGraphs" :title="graph.title" :key="graph.name">
                    <component :is="graph.component" class="graph-item"/>
                </v-graph-box>
            </div>
        </main>
    </div>
</template>

<script>
    import vToolbox from '@/components/v-toolbox.vue';
    import vGraphBox from '@/components/v-graph-box.vue';
    import vButton from '@/components/v-button.vue';
    import {mapGetters} from 'vuex';

    export default {
        name: "app",
        computed: {
            ...mapGetters(['graphs']),
            activeGraphs() {
                return this.graphs.filter(g => g.show)
            }
        },
        components: {
            vToolbox,
            vGraphBox,
            vButton,
        }
    }
</script>
<style lang="scss">
    .main-content {
        position: relative;
        display: flex;
        flex-flow: row wrap;
        align-items: flex-start;
        justify-content: center;
        align-content: flex-start;

        min-height: calc(100vh - #{$header-height} - #{$footer-height});
        .graph-item {
            margin-top: 1rem;
        }
    }
</style>