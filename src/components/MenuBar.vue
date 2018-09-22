<template>
    <div class="menu-container">
        <input type="file" ref="fileInput" @change="addFiles($event.target)" multiple
               accept="text/css, text/html"/>
        <font-awesome-icon :icon="icons.add" size="3x" class="btn btn-add-files" @click="selectFile" title="Add Files"/>
        <font-awesome-icon :icon="icons.remove" size="3x" class="btn btn-remove-files" @click="removeFiles"
                           title="Remove Files" v-show="files.length"/>
        <div class="divider-v" v-show="files.length"></div>
        <font-awesome-icon :icon="icons.analyze" size="3x" class="btn btn-analyze-file" @click="uploadFiles"
                           title="Analyze" v-show="files.length"/>
    </div>
</template>

<script>
    import {mapActions, mapGetters, mapMutations} from 'vuex';
    import {faFlask, faPlus, faTimes} from '@fortawesome/free-solid-svg-icons';

    export default {
        data() {
            return {
                icons: {
                    add: faPlus,
                    remove: faTimes,
                    analyze: faFlask,
                }
            }
        },
        computed: {
            ...mapGetters(['files'])
        },
        methods: {
            ...mapMutations(['addFiles', 'removeFiles']),
            ...mapActions(['uploadFiles']),
            selectFile() {
                this.$refs.fileInput.click()
            }
        }
    }
</script>

<style scoped>

</style>