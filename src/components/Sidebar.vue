<template>
    <div class="sidebar">
        <div v-show="!files.length" class="info">
            HTML & CSS files will show here
        </div>
        <FileList/>
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
    .info {
        font-weight: bold;
        color: lighten($dark-gray-color, 10%);
        font-size: 1.5em;
        text-align: center;
        line-height: 1.5em;
        padding: 15px 5px;
    }
</style>
