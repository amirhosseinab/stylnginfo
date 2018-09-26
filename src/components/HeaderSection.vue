<template>
    <div class="header-container">
        <div class="btn-wrapper">
            <input type="file" ref="fileInput" @change="handleFiles($event.target)" multiple
                   accept="text/css, text/html"/>

            <div class="btn add-files" @click="selectFile" v-if="!analyzed">
                <font-awesome-icon :icon="icons.add" size="3x" class="icon"
                                   :class="{'disabled':waiting}"/>
                <span>Add HTML & CSS Files</span>
            </div>
            <div class="btn remove-files" @click="renew" v-if="files.length">
                <font-awesome-icon :icon="icons.remove" size="3x" class="icon"
                                   :class="{'disabled':waiting}"/>
                <span>Remove All</span>
            </div>
            <div class="btn analyze-files" @click="analyze" v-if="files.length">
                <font-awesome-icon :icon="icons.analyze" size="3x" class="icon"
                                   :class="{'disabled':waiting}"/>
                <span>Analyze</span>
            </div>
        </div>
        <!--<div class="status-bar">-->
            <!--<div>-->
                <!--&lt;!&ndash;<span class="icon" v-show="analyzed">&ndash;&gt;-->
                    <!--&lt;!&ndash;<font-awesome-icon :icon="filter.showCssFiles ? icons.tick:icons.box" size="2x"&ndash;&gt;-->
                                       <!--&lt;!&ndash;class="check-icon"&ndash;&gt;-->
                                       <!--&lt;!&ndash;@click="filter.showCssFiles = !filter.showCssFiles"/>&ndash;&gt;-->
                <!--&lt;!&ndash;</span>&ndash;&gt;-->
                <!--&lt;!&ndash;<span v-show="analyzed">Show CSS Files</span>&ndash;&gt;-->

                <!--&lt;!&ndash;<span class="icon" v-show="analyzed">&ndash;&gt;-->
                    <!--&lt;!&ndash;<font-awesome-icon :icon="filter.showSelectors ? icons.tick:icons.box" size="2x"&ndash;&gt;-->
                                       <!--&lt;!&ndash;class="check-icon"&ndash;&gt;-->
                                       <!--&lt;!&ndash;@click="filter.showSelectors = !filter.showSelectors"/>&ndash;&gt;-->
                <!--&lt;!&ndash;</span>&ndash;&gt;-->
                <!--&lt;!&ndash;<span v-show="analyzed">Show Selectors</span>&ndash;&gt;-->
            <!--</div>-->
        <!--</div>-->
    </div>
</template>

<script>
    import {mapActions, mapGetters, mapMutations} from 'vuex';
    import {faCheckSquare, faFlask, faPlus, faSquare, faTrashAlt} from '@fortawesome/free-solid-svg-icons';

    export default {
        data() {
            return {
                icons: {
                    add: faPlus,
                    remove: faTrashAlt,
                    analyze: faFlask,
                    box: faSquare,
                    tick: faCheckSquare,
                }
            }
        },
        computed: {
            ...mapGetters(['files', 'waiting', 'analyzed'])
        },
        methods: {
            ...mapMutations(['handleFiles', 'renew']),
            ...mapActions(['analyze']),
            selectFile() {
                this.$refs.fileInput.click()
            }
        }
    }
</script>