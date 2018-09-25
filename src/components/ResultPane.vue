<template>
    <div class="result-pane">
        <Waiting/>
        <div class="result-viewer">
            <ul class="data root">
                <li v-for="hf in htmlFiles" class="data-item html">
                    <span class="title">{{hf.name}}</span>
                    <ul class="data" v-show="filter.showCssFiles">
                        <li v-for="cf in hf.cssFiles" class="data-item css">
                            <span class="title">{{cf.name}}</span>
                            <ul class="data" v-show="filter.showSelectors">
                                <li v-for="selector in cf.selectors" class="data-item selector">
                                    <span class="title">{{selector.name}}</span>
                                </li>
                            </ul>
                        </li>
                    </ul>
                </li>
            </ul>
        </div>
    </div>
</template>

<script>
    import {mapGetters} from 'vuex'
    import Waiting from "@/components/Waiting";

    export default {
        components: {
            Waiting
        },
        computed: {
            ...mapGetters(['htmlFiles', 'filter'])
        },
    }
</script>


<style lang="scss">
    .result-viewer {
        margin: 30px;
        height: calc(100vh - #{$header-height} - 60px);
        text-align: center;
        ul.data {
            margin: 0 auto 10px auto;
            display: block;
            li.data-item {
                text-align: center;
                padding: 5px 15px;
                margin: 1px;
                font-family: monospace;
                font-size: 1.2em;
                border-radius: 3px;
                display: inline-block;
                -webkit-column-break-after: avoid;
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
                    &:hover {
                        background-color: lighten($mid-gray-color, 5%);
                    }
                }
                &.css {
                    background-color: lighten($mid-gray-color, 10%);
                    color: $light-gray-color;
                    font-size: 1em;
                    &:hover {
                        background-color: lighten($mid-gray-color, 15%);
                    }
                }
                &.selector {
                    background-color: lighten($mid-gray-color, 20%);
                    color: $light-gray-color;
                    font-size: .95em;
                    &:hover {
                        background-color: lighten($mid-gray-color, 25%);
                    }
                }
            }
        }
    }

</style>