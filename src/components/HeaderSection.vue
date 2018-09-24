<template>
    <div class="header-container">
        <div class="btn-wrapper">
            <input type="file" ref="fileInput" @change="addFiles($event.target)" multiple
                   accept="text/css, text/html"/>

            <div class="btn add-files" @click="selectFile" v-show="!analyzed">
                <font-awesome-icon :icon="icons.add" size="3x" class="icon"
                                   :class="{'disabled':waiting}"/>
                <span>Add HTML & CSS Files</span>
            </div>
            <div class="btn remove-files" @click="removeFiles" v-show="files.length">
                <font-awesome-icon :icon="icons.remove" size="3x" class="icon"
                                   :class="{'disabled':waiting}"/>
                <span>Remove All Files</span>
            </div>
            <div class="btn analyze-files" @click="analyzeFiles" v-show="files.length">
                <font-awesome-icon :icon="icons.analyze" size="3x" class="icon"
                                   :class="{'disabled':waiting}"/>
                <span>Analyze</span>
            </div>
        </div>
        <div class="status-bar">
            <div>
                <span v-show="htmlFiles.length">HTML Files: {{htmlFiles.length}}</span>
                <span v-show="cssFiles.length">CSS Files: {{cssFiles.length}}</span>
            </div>
        </div>
    </div>
</template>

<script>
    import {mapActions, mapGetters, mapMutations} from 'vuex';
    import {faFlask, faPlus, faTrashAlt} from '@fortawesome/free-solid-svg-icons';

    export default {
        data() {
            return {
                icons: {
                    add: faPlus,
                    remove: faTrashAlt,
                    analyze: faFlask,
                }
            }
        },
        computed: {
            ...mapGetters(['files', 'htmlFiles', 'cssFiles', 'waiting', 'analyzed'])
        },
        methods: {
            ...mapMutations(['addFiles', 'removeFiles']),
            ...mapActions(['analyzeFiles']),
            selectFile() {
                this.$refs.fileInput.click()
            }
        }
    }
</script>

<style scoped>

</style>