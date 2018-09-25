<template>
    <div class="sidebar">
        <FileList/>
        <!--<FilterPanel/>-->
    </div>
</template>

<script>
    import {mapGetters} from 'vuex';
    import FileList from "@/components/FileList";

    export default {
        components: {
            FileList
        },
        computed: {
            ...mapGetters(['files'])
        }
    }
</script>
<style lang="scss">

</style>
