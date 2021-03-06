<template>
    <div class="scrutinize-pane">
        <v-wait-modal :show="inProgress"/>
        <div class="info-container">
            <div class="info-item">
                <div class="info-title">Index File</div>
                <div class="info-value" :class="{'missing-data':!indexFile}">
                    <font-awesome-icon :icon="indexFile?icons.ok:icons.reject"/>
                </div>
            </div>
            <div class="info-item expand">
                <div class="info-title">Total Modules</div>
                <div class="info-value html" :class="{'missing-data':!modules.length}">{{modules.length}}</div>
            </div>
            <div class="info-item">
                <div class="info-title">Total HTML Files</div>
                <div class="info-value html" :class="{'missing-data':!htmlFiles.length}">{{htmlFiles.length}}</div>
            </div>
            <div class="info-item">
                <div class="info-title">Total CSS Files</div>
                <div class="info-value" :class="{'missing-data':!cssFiles.length}">{{cssFiles.length}}</div>
            </div>
            <div class="warning" v-if="showWarning">
                Because of lacking HTML files, some graphs will be unusable.
            </div>
            <hr class="h-line" v-show="elapsedTime">
            <div class="info-item expand" v-show="elapsedTime">
                <div class="info-title">Elapsed Time</div>
                <div class="info-value">{{elapsedTime}}s</div>
            </div>
        </div>
        <v-button title="Scrutinize Now" :icon="icons.analyze" class="btn-scrutinize" :disable="isScrutinizeDisabled"
                  :class="{'disabled-button':isScrutinizeDisabled}" @click="scrutinize"/>
    </div>
</template>

<script>
    import vButton from '@/components/v-button.vue';
    import vWaitModal from '@/components/v-wait-modal.vue';
    import {faCheck, faFlask, faTimes} from '@fortawesome/free-solid-svg-icons';
    import {mapActions, mapGetters} from 'vuex';

    export default {
        name: "v-toolbox-scrutinize",
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
            ...mapGetters(['modules', 'selectedFiles', 'indexFile', 'graphs']),
            htmlFiles() {
                return this.selectedFiles.filter(sf => sf.type === "text/html")
            },
            cssFiles() {
                return this.selectedFiles.filter(sf => sf.type === "text/css")
            },
            isScrutinizeDisabled() {
                return !(this.cssFiles.length &&
                    !!this.indexFile);
            },
            showWarning() {
                return !this.isScrutinizeDisabled && this.htmlFiles.length === 0;
            },
            inProgress() {
                return this.graphs.some(g => g.inProgress)
            },
            elapsedTime() {
                return null
            }
        },
        methods: {
            ...mapActions(['getCssFilesWeight','scrutinize']),
        },
        components: {
            vButton,
            vWaitModal,
        }
    }
</script>

<style scoped lang="scss">
    .scrutinize-pane {
        position: relative;
        display: flex;
        flex-flow: column wrap;
        align-items: center;
        justify-content: space-between;
        padding: 1rem;
    }

    .info-container {
        flex: 0 1 29.5rem;
        display: flex;
        flex-flow: row wrap;

        justify-content: flex-start;
        align-content: flex-start;
        position: relative;
        align-self: stretch;
        overflow: auto;
        @extend %scrollbar;
        .info-item {
            flex: 1 0 auto;

            display: flex;
            flex-flow: row wrap;
            align-items: flex-start;
            justify-content: flex-start;
            margin: .25rem;
            font-size: .8rem;
            color: $white-color;

            transition: all .2s ease-in-out;
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
                    &.html {
                        color: $yellow-color;
                    }
                }
            }
        }
        .warning {
            flex-grow: 1;
            color: $yellow-color;
            font-style: italic;
            font-size: .75rem;
            margin: 1.2rem .2rem;
            text-align: center;
        }
    }

    .h-line {
        width: 100%;
        margin: 1rem;
        border: solid 2px darken($mid-gray-color, 10%);
        border-radius: 10px;
    }

    @media all and (max-height: $sm__height-limit) {
        .info-container {
            flex: 0 1 22.5rem;
        }
    }

    .btn-scrutinize {
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