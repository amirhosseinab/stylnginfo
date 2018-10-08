<template>
    <div>
        <main>
            <v-toolbox/>
            <div class="main-content">
                <v-graph/>
            </div>
        </main>
    </div>
</template>

<script>
    import vToolbox from '@/components/v-toolbox.vue';
    import vGraph from '@/components/v-graph.vue';

    export default {
        name: "app",
        components: {
            vToolbox,
            vGraph,
        }
    }
</script>
<style lang="scss">
    .main-content {
        display: flex;
        flex-flow: row wrap;
        align-items: flex-start;
        justify-content: center;
        align-content: flex-start;

        min-height: calc(100vh - #{$header-height} - #{$footer-height});
    }
</style>