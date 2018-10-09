<template>
    <div class="analyze-pane">
        <v-wait-modal :show="analyzing"/>
        <div class="info-container">
            <div class="info-item">
                <div class="info-title">Index File</div>
                <div class="info-value" :class="{'missing-data':!indexFile}">
                    <font-awesome-icon :icon="indexFile?icons.ok:icons.reject"/>
                </div>
            </div>
            <div class="info-item expand">
                <div class="info-title">Total Modules</div>
                <div class="info-value" :class="{'missing-data':!modules.length}">{{modules.length}}</div>
            </div>
            <div class="info-item">
                <div class="info-title">Total HTML Files</div>
                <div class="info-value" :class="{'missing-data':!htmlFiles.length}">{{htmlFiles.length}}</div>
            </div>
            <div class="info-item">
                <div class="info-title">Total CSS Files</div>
                <div class="info-value" :class="{'missing-data':!cssFiles.length}">{{cssFiles.length}}</div>
            </div>
            <div class="info-item expand" v-if="analyzingTime">
                <div class="info-title">Elapsed Time</div>
                <div class="info-value">{{analyzingTime}}s</div>
            </div>
        </div>
        <v-button title="Analyze Now" :icon="icons.analyze" class="btn-analyze" :disable="isAnalyzedDisabled"
                  :class="{'disabled-button':isAnalyzedDisabled}" @click="analyze"/>
    </div>
</template>

<script>
    import vButton from '@/components/v-button.vue';
    import vWaitModal from '@/components/v-wait-modal.vue';
    import {faCheck, faFlask, faTimes} from '@fortawesome/free-solid-svg-icons';
    import {mapActions, mapGetters} from 'vuex';

    export default {
        name: "v-toolbox-analyze",
        data() {
            return {
                icons: {
                    analyze: faFlask,
                    ok: faCheck,
                    reject: faTimes,
                }
            }
        },
        computed: {
            ...mapGetters(['modules', 'selectedFiles', 'indexFile', 'analyzing', 'analyzingTime']),
            htmlFiles() {
                return this.selectedFiles.filter(sf => sf.type === "text/html")
            },
            cssFiles() {
                return this.selectedFiles.filter(sf => sf.type === "text/css")
            },
            isAnalyzedDisabled() {
                return !(this.modules.length &&
                    this.cssFiles.length &&
                    this.htmlFiles.length &&
                    !!this.indexFile);
            }
        },
        methods: {
            ...mapActions(['analyze']),
        },
        components: {
            vButton,
            vWaitModal,
        }
    }
</script>

<style scoped lang="scss">
    .analyze-pane {
        position: relative;
        display: flex;
        flex-flow: column wrap;
        align-items: center;
        justify-content: space-between;
        padding: 1rem;
    }

    .info-container {
        flex: 0 1 27rem;
        display: flex;
        flex-flow: row wrap;

        justify-content: flex-start;
        align-content: flex-start;
        padding: .5rem;
        position: relative;
        align-self: stretch;
        overflow: auto;
        @extend %scrollbar;
        .info-item {
            flex: 1 0 10rem;

            display: flex;
            flex-flow: row wrap;
            align-items: flex-start;
            justify-content: flex-start;
            margin: .25rem;
            font-size: .8rem;
            color: $white-color;
            &.expand {
                flex-basis: 20rem;
            }

            .info-title, .info-value {
                flex: auto 0;
                display: flex;
                align-items: center;
                min-height: 2.5rem;
                background-color: darken($mid-gray-color, 15%);
            }
            .info-title {
                flex: 1 0 4rem;
                align-self: center;
                padding: .5rem .4rem .5rem .8rem;
                border-radius: 20px 0 0 20px;
                margin-right: .2rem;
                font-size: .75rem;
            }
            .info-value {
                flex: 0 0 3.5rem;
                justify-content: center;
                padding: .5rem .8rem .5rem .5rem;
                border-radius: 0 20px 20px 0;
                font-weight: bold;
                font-size: .8rem;
                color: $green-color;
                &.missing-data {
                    color: $light-red-color;
                }
            }
        }
    }

    @media all and (max-height: $sm__height-limit) {
        .info-container {
            flex: 0 1 23.5rem;
        }
    }

    .btn-analyze {
        flex: 0 1 auto;
        margin: 1.5rem;
        width: 11rem;
        font-size: .9rem;
        font-weight: bold;
        height: 3rem;
        background-color: $orange-color;
        color: $dark-gray-color;
        border-color: $dark-gray-color;
        &.disabled-button {
            background-color: lighten($mid-gray-color, 10%);
            color: lighten($dark-gray-color, 8%);
            border-color: lighten($dark-gray-color, 8%);
            box-shadow: none;
            cursor: not-allowed;
        }
    }

</style>