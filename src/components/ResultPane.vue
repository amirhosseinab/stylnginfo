<template>
    <div class="result-pane">
        <Waiting/>
        <div class="result-viewer">
            <ul class="data root">
                <li v-for="hf in htmlFiles" class="data-item html">
                    <span class="title">{{hf.name}}</span>
                    <ul class="data">
                        <li v-for="cf in hf.cssFiles" class="data-item css">
                            <span class="title">{{cf.name}}</span>
                            <ul class="data">
                                <li v-for="selector in cf.selectors" class="data-item selector">
                                    <span class="title">{{selector.name}}</span>
                                </li>
                            </ul>
                        </li>
                    </ul>
                </li>
            </ul>
        </div>
        <!--<div v-for="hf in htmlFileNames" class="html-file">-->
        <!--<div>{{hf.name}}</div>-->
        <!--<div class="css-files-block">-->
        <!--<div v-for="cf in hf.cssFiles" class="css-file">-->
        <!--{{cf.name}}-->
        <!--<div class="selectors-block">-->
        <!--<div v-for="s in cf.selectors" class="selector"-->
        <!--@click="selectSelectors({'htmlFileName':hf.name,'selector':s.name})"-->
        <!--:class="{'selected':s.selected}">-->
        <!--{{s.name}}-->
        <!--</div>-->
        <!--</div>-->
        <!--</div>-->
        <!--</div>-->
        <!--</div>-->
    </div>
</template>

<script>
    import {mapGetters} from 'vuex'
    import Waiting from "@/components/Waiting";

    export default {
        data() {
            return {
                data: [10, 71, 78, 25, 36, 92, 20],
                line: '',
            }
        },
        components: {
            Waiting
        },
        computed: {
            ...mapGetters(['htmlFiles'])
        },
    }
</script>


<style lang="scss">
    .result-viewer {
        margin: 20px;

        height: calc(100vh - #{$header-height} - 40px);
        text-align: center;
        @extend %mid-align;
        ul.data {
            display: block;
            vertical-align: middle;
            &.root {
                display: inline-block;
            }
            li.data-item {
                text-align: center;
                padding: 5px 15px;
                margin: 1px;
                font-family: monospace;
                font-size: 1.2em;
                border-radius: 3px;
                display: inline-block;
                max-width: 400px;
                vertical-align: top;
                cursor: pointer;
                .title {
                    display: inline-block;
                    padding: 3px 0 7px 0;
                }
                &.html {
                    background-color: $mid-gray-color;
                    color: $light-gray-color;
                    &:hover{
                        background-color: lighten($mid-gray-color,2%);
                    }
                }
                &.css {
                    background-color: lighten($mid-gray-color, 5%);
                    color: $light-gray-color;
                    font-size: 1em;
                    &:hover{
                        background-color: lighten($mid-gray-color,7%);
                    }
                }
                &.selector {
                    background-color: lighten($mid-gray-color, 10%);
                    color: $light-gray-color;
                    font-size: .95em;
                    &:hover{
                        background-color: lighten($mid-gray-color,12%);
                    }
                }
            }
        }
    }

</style>