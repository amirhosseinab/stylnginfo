<template>
    <div class="main-panel">
        <Waiting/>
        <Graph-visualization v-if="analyzed"/>
    </div>
</template>

<script>
    import Waiting from "@/components/Waiting";
    import GraphVisualization from '@/components/GraphVisualization';
    import {mapGetters} from 'vuex'

    export default {
        components: {
            Waiting,
            GraphVisualization
        },
        computed: {
            ...mapGetters(['analyzed']),
        }
    }
</script>