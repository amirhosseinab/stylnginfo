<template>
    <div class="file-list-pane">
        <v-wait-modal :show="inProgress"/>

        <input type="file" ref="filesInput" @change="addSelectedFiles($refs.filesInput)" multiple
               accept="text/css, text/html, .less, .scss"/>
        <input type="file" ref="indexFileInput" @change="addIndexFile($refs.indexFileInput)" accept="text/html"/>

        <div class="btn-group">
            <div class="btn-group-add">
                <v-button title="Add Files" :icon="icons.addFiles" @click="selectFiles" class="btn-add-files btn"/>
                <v-button v-if="!indexFile" title="Add Index File" :icon="icons.addIndexFile" @click="selectIndexFile"
                          class="btn"/>
                <div v-else class="index-file">
                    <span class="file-name">{{indexFile.name}}</span>
                    <font-awesome-icon :icon="icons.remove" class="remove-button" @click="removeIndexFile"/>
                </div>
            </div>


            <v-button title="Remove All" :icon="icons.remove" @click="removeAllFiles" class="btn-remove-files btn"
                      v-if="selectedFiles.length"/>
        </div>
        <div class="module-list">
            <div class="info module-list" v-if="!modules.length">HTML Modules</div>
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
                <font-awesome-icon :icon="icons.cssFile" class="css-file-icon" size="2x"/>
                <div>{{file.name}}</div>
                <font-awesome-icon :icon="icons.remove" class="remove-button" @click="removeFile(file)"/>
            </div>
            <div v-for="file in scssFiles" :key="file.name" class="file-item scss file-name">
                <font-awesome-icon :icon="icons.scssFile" class="scss-file-icon" size="2x"/>
                <div>{{file.name}}</div>
                <font-awesome-icon :icon="icons.remove" class="remove-button" @click="removeFile(file)"/>
            </div>
            <div v-for="file in lessFiles" :key="file.name" class="file-item less file-name">
                <font-awesome-icon :icon="icons.lessFile" class="less-file-icon" size="2x"/>
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
    import {faCheck, faFileAlt, faPlus, faTimes,} from "@fortawesome/free-solid-svg-icons";
    import {faCss3, faLess, faSass} from "@fortawesome/free-brands-svg-icons";

    export default {
        name: "v-toolbox-files",
        data() {
            return {
                icons: {
                    addFiles: faPlus,
                    addIndexFile: faFileAlt,
                    remove: faTimes,
                    accept: faCheck,
                    lessFile: faLess,
                    scssFile: faSass,
                    cssFile: faCss3,
                },
            }
        },
        components: {
            vButton,
            vWaitModal,
        },
        computed: {
            ...mapGetters(['selectedFiles', 'indexFile', 'modules', 'graphs']),

            htmlFiles() {
                return this.selectedFiles.filter(f => f.type === "text/html")
            },
            cssFiles() {
                return this.selectedFiles.filter(f => f.type === "text/css")
            },
            lessFiles() {
                return this.selectedFiles.filter(f => f.name.endsWith(".less"))
            },
            scssFiles() {
                return this.selectedFiles.filter(f => f.name.endsWith(".scss"))
            },
            inProgress() {
                return this.graphs.some(g => g.inProgress)
            },
        },
        methods: {
            ...mapMutations(['addSelectedFiles', 'addIndexFile', 'removeIndexFile', 'removeFile', 'removeAllFiles', 'addModule', 'removeModule']),
            selectFiles() {
                this.$refs.filesInput.click()
            },
            selectIndexFile() {
                this.$refs.indexFileInput.click();
            },
            updateModule(module) {
                if (module.name !== "") {
                    module.editMode = false;
                }
            },
        }
    }
</script>

<style scoped lang="scss">
    $file-list-background-color: darken($mid-gray-color, 7%);
    $info-color: lighten($file-list-background-color, 5%);
    $button-group-height: 3.5rem;
    $module-list-height: 3.5rem;
    $body-content-height: calc(#{$toolbox-body-height} - #{$toolbox-body-border-bottom-width});

    .file-list-pane {
        display: flex;
        flex-flow: column wrap;
        align-items: stretch;
        justify-content: stretch;
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
        .btn {
            background-color: lighten($mid-gray-color, 3%);
        }
        .btn-group-add {
            flex: 0 1 64%;
            display: flex;
            flex-flow: row nowrap;
            align-items: center;
            justify-content: space-between;
            .btn-add-files {
                flex: 0 0 auto;
                background-color: $dark-green-color;
            }
            .index-file {
                flex: 0 1 8rem;
                background-color: $dark-gray-color;
                padding: .5rem 1rem;
                border-radius: 8px;
                display: flex;
                align-items: center;
                justify-content: space-around;
                color: $yellow-color;
            }
        }

        .btn-remove-files {
            flex: 0 0 30%;
            background-color: $red-color;
        }

    }

    .module-list {
        flex: 0 1 $module-list-height;
        //max-height: min-content;
        display: flex;
        flex-flow: row wrap;
        align-items: flex-start;
        justify-content: flex-start;
        overflow: auto;
        padding: .2rem;
        background-color: $file-list-background-color;
        margin: 0 .5rem;
        border-radius: 8px;
        @extend %scrollbar;
        &.info {
            color: $info-color;
            font-size: 1.7rem;
            align-self: center;
            justify-content: center;
        }
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

        flex: 1 0 calc(#{$body-content-height} - #{$button-group-height} - #{$module-list-height} - 1rem);
        overflow: auto;
        transform: translateY(.5rem);
        margin: 0 .5rem;
        border-radius: 8px;

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
            &.css, &.less, &.scss {
                padding: 0 .5rem 0 .3rem;
                font-weight: bold;
                color: $green-color;
                border-radius: 8px;
                background-color: lighten($dark-gray-color, 12%);
            }
            &.css {
                .css-file-icon {
                    padding: .2rem;
                    margin-right: .1rem;
                    color: $orange-color;
                }
            }
            &.scss {
                .scss-file-icon {
                    padding: .1rem;
                    margin: 0 .3rem 0 0;
                    color: $purple-color;
                }
            }
            &.less {
                .less-file-icon {
                    margin: 0 .3rem 0 0;
                    color: $blue-color;
                }
            }
        }
    }

    @media all and (max-height: $sm__height-limit) {
        $body-content-height: calc(#{$toolbox-body-height__sm} - #{$toolbox-body-border-bottom-width});
        .file-list {
            flex: 1 0 calc(#{$body-content-height} - #{$button-group-height} - #{$module-list-height} - 1rem);
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
        margin: 0 .5rem;
        align-self: center;
        color: $info-color;
        font-weight: bold;
        font-size: 2.5rem;
    }
</style>