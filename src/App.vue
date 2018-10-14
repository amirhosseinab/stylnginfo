<template>
    <div>
        <main>
            <v-toolbox/>
            <div class="main-content">
                <v-user-guide v-if="!selectedGraph"/>
                <v-graph-box v-else :title="selectedGraph.title">
                    <component slot="graph" :is="selectedGraph.component" class="graph-item" :data="selectedGraph.data"/>
                </v-graph-box>
            </div>
        </main>
    </div>
</template>

<script>
    import vToolbox from '@/components/v-toolbox.vue';
    import vGraphBox from '@/components/v-graph-box.vue';
    import vUserGuide from '@/components/v-user-guide.vue';
    import vGraphCssFilesWeight from '@/components/v-graph-css-files-weight.vue';

    import {mapGetters} from 'vuex';

    export default {
        name: "app",
        computed: {
            ...mapGetters(['graphs', 'selectedGraph']),
        },
        components: {
            vToolbox,
            vGraphBox,
            vUserGuide,
            vGraphCssFilesWeight,
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
            margin-top: .5rem;
        }
    }
</style>