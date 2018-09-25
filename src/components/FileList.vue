<template>
    <div class="file-list-container">
        <div class="toolbar file-list">
            <span v-show="analyzed">
                <font-awesome-icon :icon="allSelected? icons.tick:icons.box" size="2x" class="toolbar toggle-selector"
                                   @click="toggleAllFiles"/>
                <span class="title">Toggle Selection</span>
            </span>
        </div>
        <div v-show="!files.length" class="info">
            HTML & CSS files will show here
        </div>

        <ul class="file-list">
            <li v-for="f in files" class="file-list-item"
                :class="{'html':f.file.type==='text/html','css':f.file.type==='text/css', 'permanent':analyzed}">

                <font-awesome-icon :icon="f.selected?icons.tick:icons.box" size="2x" class="icon checkbox"
                                   @click="toggleFileSelection(f)" v-show="analyzed"/>

                <span class="file-name">{{f.file.name}}</span>
                <font-awesome-icon :icon="icons.remove" class="icon remove" size="2x" @click="removeFile(f)"
                                   :class="{'disabled':waiting}" v-show="!analyzed"/>
            </li>
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
            ...mapGetters(['files', 'analyzed', 'waiting']),
            allSelected: function () {
                return this.files.every(f => f.selected);
            },

        },
        methods: {
            ...mapMutations(['removeFile', 'toggleFileSelection', 'toggleAllFiles'])
        }
    }
</script>
<style lang="scss">
    .file-list-container {
        height: 100%;
    }

    .toolbar.file-list {
        display: block;
        height: 40px;
        background-color: $mid-gray-color;
        border-top: solid 1px $border-color;
        @extend %mid-align;
        .toolbar.toggle-selector {
            vertical-align: middle;
            padding: 4px;
            margin-left: 4px;
            cursor: pointer;
        }
        span {
            display: inline-block;
            vertical-align: middle;
            &.title {
                padding: 2px 5px 2px 0;
                font-size: .9em;
                font-weight: bold;
            }
        }

    }

    ul.file-list {
        display: block;
        margin: 0;
        padding: 2px 0;
        text-align: left;
        height: calc(100% - 40px);
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