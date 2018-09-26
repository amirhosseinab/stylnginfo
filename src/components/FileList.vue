<template>
    <div class="file-list-container">
        <div class="toolbar file-list">
            <div v-show="analyzed" class="toolbar-item toggle-selector" @click="toggleAllFilesSelection">
                <font-awesome-icon :icon="allSelected? icons.tick:icons.box" size="2x" class="toggle-selector"/>
                <span class="title">Toggle Selection</span>
            </div>
        </div>

        <div v-show="!files.length" class="info">
            HTML & CSS files will show here
        </div>

        <ul class="file-list">
            <template v-if="!analyzed">
                <li v-for="f in files" class="file-list-item"
                    :class="{'html':f.type==='text/html','css':f.type==='text/css'}">
                    <span class="file-name">{{f.name}}</span>
                    <font-awesome-icon :icon="icons.remove" class="icon remove" size="2x" @click="removeFile(f)"
                                       :class="{'disabled':waiting}" v-if="!analyzed"/>
                </li>
            </template>
            <template v-else>
                <li v-for="f in htmlFiles.concat(cssFiles)" class="file-list-item permanent"
                    :class="{'html':f.fileType === 0,'css':f.fileType === 1}" @click="toggleFileSelection(f)">
                    <font-awesome-icon :icon="f.selected?icons.tick:icons.box" size="2x" class="icon checkbox"/>
                    <span class="file-name">{{f.name}}</span>
                </li>
            </template>
        </ul>
    </div>
</template>
<script>
    import {faCheckSquare, faFile, faSquare, faTimes} from '@fortawesome/free-solid-svg-icons';
    import {mapGetters, mapMutations} from 'vuex';

    export default {
        data() {
            return {
                icons: {
                    remove: faTimes,
                    file: faFile,
                    box: faSquare,
                    tick: faCheckSquare,
                }
            }
        },
        computed: {
            ...mapGetters(['files', 'htmlFiles', 'cssFiles', 'analyzed', 'waiting']),
            allSelected: function () {
                return this.htmlFiles.concat(this.cssFiles).every(f => f.selected);
            },

        },
        methods: {
            ...mapMutations(['removeFile', 'toggleFileSelection', 'toggleAllFilesSelection'])
        }
    }
</script>
<style lang="scss">
    .toolbar {
        .toolbar-item {
            &.toggle-selector {
                cursor: pointer;
            }
        }
        &.file-list {
            .toggle-selector {
                vertical-align: middle;
                padding: 4px;
                cursor: pointer;
            }
        }
    }

    ul.file-list {
        display: block;
        margin: 0;
        padding: 2px 0;
        text-align: left;
        height: calc(100% - #{$toolbar-height});
        overflow: auto;
        @extend %scrollbar;

        li.file-list-item {
            @extend %mid-align;
            font-family: monospace;
            overflow: hidden;
            height: 25px;
            font-size: 1.1em;
            text-align: left;
            margin: 3px;
            vertical-align: middle;
            display: inline-block;
            background-clip: padding-box;
            border-radius: 20px 50px 50px 20px;
            background-color: $blue-color;
            padding-left: 4px;
            cursor: default;
            &.permanent {
                padding-right: 10px;
                padding-left: 0;
                cursor: pointer;
            }
            .icon {
                width: 25px;
                height: 25px;
                padding: 3px;
                display: inline-block;
                text-align: center;
                vertical-align: middle;
                cursor: pointer;
            }
            .file-name {
                color: $light-gray-color;
                padding: 4px 1px 5px 5px;
                display: inline-block;
                vertical-align: middle;
            }
            &.html {
                background-color: $blue-color;
            }
            &.css {
                background-color: $green-color;
            }
            .icon.checkbox {
                border-radius: 3px 0 0 3px;
                background-color: $mid-gray-color;
                color: $light-gray-color;
            }
            .remove.icon {
                color: $dark-gray-color;
                padding: 6px;
                border-radius: 50px;
            }
        }
    }

    .info {
        font-weight: bold;
        color: lighten($dark-gray-color, 10%);
        font-size: 1.5em;
        width: $sidebar-width;
        text-align: center;
        line-height: 1.5em;
        padding: 15px 5px;
        position: absolute;
    }
</style>