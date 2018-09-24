<template>
    <div class="filter-panel">
        <div class="filter-toolbox">

        </div>
        <ul>
            <li v-for="f in htmlFileNames" class="filter-item">
                <span class="html" :class="{'selected':false}">{{f}}</span>
            </li>
            <li v-for="f in cssFileNames" class="filter-item">
                <span class="css" :class="{'selected':false}">{{f}}</span>
            </li>
        </ul>
    </div>
</template>
<script>
    import {mapGetters} from 'vuex';

    export default {
        computed: {
            ...mapGetters(['htmlFileNames', 'cssFileNames'])
        }
    }
</script>
<style lang="scss">
    .filter-panel {
        max-height: calc(100% - #{$file-list-height});
        ul {
            overflow: auto;
            margin: 0;
            padding: 5px;
            display: block;
            height: calc(100%  - 30px);
            @extend %scrollbar;

            .filter-item {
                display: inline-block;
                margin: 2px;
                cursor: pointer;
                .html, .css {
                    padding: 3px 7px;
                    border-radius: 20px;
                    font-size: 0.9em;
                    display: inline-block;
                    color: $dark-gray-color;
                    &.selected {
                        background-color: $orange-color !important;
                    }
                }
                .html {
                    background-color: $blue-color;
                }
                .css {
                    background-color: $green-color;
                }

            }
        }
        .filter-toolbox {
            height: 30px;
            background-color: $mid-gray-color;
        }
    }
</style>