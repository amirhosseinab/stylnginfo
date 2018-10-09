<template>
    <div class="file-list-pane">
        <v-wait-modal :show="analyzing"/>

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
        <div class="module-list" v-if="modules.length">
            <div v-for="(module,index) in modules" :key="index" class="module-item"
                 :class="{'edit-mode':module.editMode}"
                 :style="{backgroundColor:module.color}">

                <div v-if="!module.editMode" class="module-text" @click="module.editMode=true">{{module.name}}</div>
                <input v-else v-model="module.name" @focus="$event.target.select()" autofocus
                       @keypress.enter="updateModule(module)" type="text" class="module-text-edit"
                       placeholder="Module Name">

                <font-awesome-icon v-if="!module.editMode" class="remove-button" :icon="icons.remove"
                                   @click="removeModule(module)"/>
                <font-awesome-icon v-else class="accept-button" :icon="icons.accept"
                                   @click="updateModule(module)"/>
            </div>
        </div>
        <div class="file-list">
            <div class="info" v-if="!selectedFiles.length">HTML & CSS Files will be Shown Here</div>
            <div v-for="file in htmlFiles" :key="file.name" class="file-item html file-name"
                 :style="{borderColor: file.module.color}">
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
    import vWaitModal from "@/components/v-wait-modal.vue";
    import {faCheck, faFileAlt, faPlus, faTimes} from "@fortawesome/free-solid-svg-icons"

    export default {
        name: "v-toolbox-files",
        data() {
            return {
                icons: {
                    addFiles: faPlus,
                    addIndexFile: faFileAlt,
                    remove: faTimes,
                    accept: faCheck,
                },
            }
        },
        components: {
            vButton,
            vWaitModal,
        },
        computed: {
            ...mapGetters(['selectedFiles', 'indexFile', 'modules', 'analyzing']),

            htmlFiles() {
                return this.selectedFiles.filter(f => f.type === "text/html")
            },
            cssFiles() {
                return this.selectedFiles.filter(f => f.type === "text/css")
            },
        },
        methods: {
            ...mapMutations(['addSelectedFiles', 'addIndexFile', 'removeIndexFile', 'removeFile', 'removeAllFiles', 'addModule', 'removeModule']),
            selectFiles() {
                this.$refs.filesInput.click()
            },
            selectIndexFile() {
                this.$refs.indexFileInput.click()
            },
            updateModule(module) {
                if (module.name !== "") {
                    module.editMode = false;
                }
            }
        }
    }
</script>

<style scoped lang="scss">
    $file-list-background-color: lighten($mid-gray-color, 2%);
    $button-group-height: 3.5rem;
    $module-list-height: 3.6rem;
    $body-content-height: calc(#{$toolbox-body-height} - #{$toolbox-body-border-bottom-width});

    .file-list-pane {
        display: flex;
        flex-flow: column wrap;
        align-items: stretch;
        justify-content: flex-start;
        position: relative;
        z-index: 99;

    }

    .btn-group {
        flex: 0 1 $button-group-height;
        display: flex;
        flex-flow: row wrap;
        justify-content: space-between;
        align-items: center;
        padding: .5rem;

        .btn-add-files, .btn-remove-files {
            flex: 0 1 6.5rem;
        }

        .index-file {
            flex: 0 1 9rem;
            background-color: $dark-gray-color;
            padding: .5rem 1rem;
            border-radius: 8px;
            display: flex;
            align-items: center;
            justify-content: space-around;
            color: $yellow-color;
        }
    }

    .module-list {
        flex: 0 1 $module-list-height;
        max-height: min-content;
        display: flex;
        flex-flow: row wrap;
        align-items: center;
        justify-content: flex-start;
        overflow: auto;
        padding: .4rem .2rem;
        background-color: darken($file-list-background-color, 14%);
        @extend %scrollbar;
        .module-item {
            display: flex;
            align-items: center;
            justify-content: space-between;
            padding: .2rem .5rem .35rem;
            margin: .1rem;
            font-size: .7rem;
            border-radius: 5px;
            color: $dark-gray-color;
            &.edit-mode {
                border-radius: 5px;
                padding: 0 .5rem .15rem .2rem;

            }
            .module-text {
                color: $dark-gray-color;
                transform: translateY(.05rem);
                &:hover {
                    cursor: pointer;
                    color: $white-color;
                }
            }
            .module-text-edit {
                outline: none;
                line-height: .5rem;
                height: 1.05rem;
                margin: .1rem 0 0;
                border-radius: 5px;
                max-width: 5rem;
                background-color: transparent;
                //border: solid 1px lighten($mid-gray-color, 5%);
                border: none;
            }
            .remove-button {
                color: $dark-gray-color;
                transform: translateY(.07rem);
                margin-left: .3rem;
            }
        }
    }

    .file-list {
        flex: 1 0 calc(#{$body-content-height} - #{$button-group-height} - #{$module-list-height});
        overflow: auto;

        display: flex;
        flex-flow: row wrap;
        align-items: flex-start;
        justify-content: flex-start;
        align-content: flex-start;

        padding: .3rem .2rem;
        background-color: $file-list-background-color;

        @extend %scrollbar;
        .file-item {
            display: flex;
            flex-flow: nowrap;
            align-items: center;

            padding: .4rem .5rem;
            margin: .15rem .2rem;
            background-color: $dark-gray-color;
            border-radius: 0 8px 8px 0;

            &.html {
                color: $blue-color;
                border-left: solid .5rem;
            }
            &.css {
                color: $green-color;
                border-radius: 8px;
                background-color: lighten($dark-gray-color, 4%);
            }
        }
    }

    @media all and (max-height: $sm__height-limit) {
        $body-content-height: calc(#{$toolbox-body-height__sm} - #{$toolbox-body-border-bottom-width});
        .file-list {
            flex: 1 0 calc(#{$body-content-height} - #{$button-group-height} - #{$module-list-height});
        }
    }

    .remove-button {
        cursor: pointer;
        margin-left: .6rem;
        padding: .01rem;
        color: $light-red-color;
        min-width: 0.5rem;
    }

    .accept-button {
        cursor: pointer;
        margin-left: .3rem;
        padding: .01rem;
        transform: translateY(.1rem);
        color: $dark-gray-color;
        min-width: 0.5rem;
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