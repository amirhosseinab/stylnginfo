<template>
    <div class="settings-container">
        <v-button title="Save Current Data" class="btn btn-save" :icon="icons.save"
                  :disable="!graphDataExists" @click="storeData"
                  :class="{'disabled':!graphDataExists}"/>

        <v-button title="Load Data" class="btn btn-load" :icon="icons.load"
                  :disable="!storedDataExists" @click="loadStoredData"
                  :class="{'disabled':!storedDataExists}"/>

        <v-button title="Remove Stored Data" class="btn btn-clear" :icon="icons.remove"
                  :disable="!storedDataExists" @click="clearStoredData"
                  :class="{'disabled':!storedDataExists}"/>
    </div>
</template>

<script>
    import vButton from "@/components/v-button.vue";
    import {faDownload, faTrashAlt, faUpload} from "@fortawesome/free-solid-svg-icons";
    import {mapGetters, mapMutations} from "vuex";

    export default {
        name: "v-toolbox-filter",
        data() {
            return {
                icons: {
                    save: faDownload,
                    load: faUpload,
                    remove: faTrashAlt,
                },
                storedDataExists: false,
            }
        },
        computed: {
            ...mapGetters(['graphs']),
            graphDataExists() {
                return this.graphs.some(g => Object.keys(g.data).length > 0)
            },
        },
        methods: {
            ...mapMutations(['removeAllFiles', 'updateGraph']),
            storeData() {
                localStorage.setItem("graphs", JSON.stringify(this.graphs))
                this.storedDataExists = true;
            },
            clearStoredData() {
                localStorage.removeItem("graphs")
                this.storedDataExists = false;
            },
            loadStoredData() {
                this.removeAllFiles();
                JSON.parse(localStorage.getItem("graphs"))
                    .forEach(g => this.updateGraph({
                        graphName: g.name,
                        data: g.data,
                        elapsedTime: null
                    }))
            },
        },
        components: {
            vButton
        },
        mounted() {
            this.storedDataExists = localStorage.getItem("graphs") !== null;
        }
    }
</script>

<style scoped lang="scss">
    .settings-container {
        display: flex;
        flex-flow: row wrap;
        align-items: center;
        justify-content: space-evenly;
        justify-items: stretch;
        padding: 1rem;

        .btn {
            flex: 1 0 8rem;
            min-height: 2.5rem;
            margin: .5rem;
            &.disabled {
                cursor: default;
                background-color: $mid-gray-color;

            }
        }
        .btn-save {
            background-color: $dark-green-color;
        }
        .btn-load {
            background-color: $dark-purple-color;
        }
        .btn-clear {
            background-color: $red-color;
        }
    }
</style>