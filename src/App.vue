<template>
    <div>
        <main>
            <v-toolbox/>
            <div class="main-content">
                <v-graph-box title="Graph No.1">
                </v-graph-box>

                <v-graph-box title="Graph No.2">
                </v-graph-box>
            </div>
        </main>
    </div>
</template>

<script>
    import vToolbox from '@/components/v-toolbox.vue';
    import vGraphBox from '@/components/v-graph-box.vue';

    export default {
        name: "app",
        components: {
            vToolbox,
            vGraphBox,
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