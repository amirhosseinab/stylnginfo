<template>
    <div class="sidebar">
        <div v-show="!files.length" class="info">
            HTML & CSS files will show here
        </div>

        <FileList/>
        <FilterPanel/>
        <div v-show="!files.length" class="info">
            Filter data will show here
        </div>
    </div>
</template>

<script>
    import {mapGetters} from 'vuex';
    import FileList from "@/components/FileList";
    import FilterPanel from "@/components/FilterPanel";

    export default {
        components: {
            FilterPanel,
            FileList
        },
        computed: {
            ...mapGetters(['files'])
        }
    }
</script>
<style lang="scss">
    .info {
        font-weight: bold;
        color: lighten($dark-gray-color, 10%);
        font-size: 1.5em;
        text-align: center;
        line-height: 1.5em;
        padding: 15px 5px;
    }
</style>
