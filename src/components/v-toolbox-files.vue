<template>
    <div class="file-list-panel">
        <input type="file" ref="filesInput" @change="addSelectedFiles($refs.filesInput)" multiple
               accept="text/css, text/html"/>
        <input type="file" ref="indexFileInput" @change="addIndexFile($refs.indexFileInput)" accept="text/html"/>

        <div class="btn-group">
            <v-button title="Add Files" :icon="icons.addFiles" @click="selectFiles" class="btn-add-files"/>
            <v-button title="Remove All" :icon="icons.remove" @click="removeAllFiles" class="btn-remove-files"
                      v-if="selectedFiles.length"/>
            <v-button v-if="!indexFile" title="Add Index File" :icon="icons.addIndexFile" @click="selectIndexFile"/>
            <div v-else class="index-file">
                <span class="file-name">{{indexFile.name}}</span>
                <font-awesome-icon :icon="icons.remove" class="remove-button" @click="removeIndexFile"/>
            </div>
        </div>

        <div class="file-list">
            <div class="info" v-if="!selectedFiles.length">HTML & CSS Files will be Shown Here</div>
            <div v-for="file in htmlFiles" :key="file.name" class="file-item html file-name">
                <div>{{file.name}}</div>
                <font-awesome-icon :icon="icons.remove" class="remove-button" @click="removeFile(file)"/>
            </div>
            <div v-for="file in cssFiles" :key="file.name" class="file-item css file-name">
                <div>{{file.name}}</div>
                <font-awesome-icon :icon="icons.remove" class="remove-button" @click="removeFile(file)"/>
            </div>
        </div>
    </div>
</template>

<script>
    import {mapGetters, mapMutations} from 'vuex';

    import vButton from "@/components/v-button.vue";
    import {faFileAlt, faPlus, faTimes} from "@fortawesome/free-solid-svg-icons"

    export default {
        name: "v-toolbox-files",
        data() {
            return {
                icons: {
                    addFiles: faPlus,
                    addIndexFile: faFileAlt,
                    remove: faTimes,
                }
            }
        },
        components: {
            vButton,
        },
        computed: {
            ...mapGetters(['selectedFiles', 'indexFile']),

            htmlFiles() {
                return this.selectedFiles.filter(f => f.type === "text/html")
            },
            cssFiles() {
                return this.selectedFiles.filter(f => f.type === "text/css")
            },
        },
        methods: {
            ...mapMutations(['addSelectedFiles', 'addIndexFile', 'removeIndexFile', 'removeFile', 'removeAllFiles']),
            selectFiles() {
                this.$refs.filesInput.click()
            },
            selectIndexFile() {
                this.$refs.indexFileInput.click()
            },
        }
    }
</script>

<style scoped lang="scss">
    $file-list-background-color: lighten($mid-gray-color, 2%);
    $body-content: calc(#{$toolbox-body-height} - #{$toolbox-body-border-bottom-width});
    $button-group-height: 3.5rem;
    .file-list-panel {
        display: flex;
        flex-flow: column wrap;
        position: relative;
    }

    .btn-group {
        display: flex;
        justify-content: space-between;
        align-items: center;
        padding: .5rem;
        max-height: $button-group-height;

        .btn-add-files, .btn-remove-files {
            min-width: 6.5rem;
        }

        .index-file {
            background-color: $dark-gray-color;
            padding: .5rem 1rem;
            border-radius: 8px;
            display: flex;
            align-items: center;
            justify-content: space-around;
            color: $yellow-color;
        }
    }

    .file-list {
        flex-grow: 1;

        display: flex;
        flex-flow: row wrap;
        align-items: flex-start;
        justify-content: flex-start;
        align-content: flex-start;

        padding: .3rem .2rem;
        overflow: auto;
        background-color: $file-list-background-color;
        height: auto;
        max-height: calc(#{$body-content} - #{$button-group-height});

        @extend %scrollbar;
        .file-item {
            display: flex;
            flex-flow: nowrap;
            align-items: center;

            padding: .4rem .7rem;
            margin: .12rem;
            background-color: $dark-gray-color;
            border-radius: 8px;

            &.html {
                color: $blue-color;
            }
            &.css {
                color: $green-color;
            }
        }
    }

    .remove-button {
        cursor: pointer;
        margin-left: .6rem;
        padding: .01rem;
        color: $light-red-color;
    }

    .info {
        flex-grow: 1;
        text-align: center;
        align-self: center;
        color: darken($file-list-background-color, 4%);
        font-weight: bold;
        font-size: 2.5rem;
    }
</style>