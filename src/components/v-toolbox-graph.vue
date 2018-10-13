<template>
    <div class="graph-selector-container">
        <div v-for="graph in graphs" :key="graph.name" class="graph-selector"
             :class="{selected:selectedGraphName===graph.name}"
             @click="updateSelectedGraph(graph.name)">
            <div class="select-icon" :class="{selected:selectedGraphName===graph.name}">
                <font-awesome-icon :icon="icons.select"/>
            </div>
            <div class="graph-data">
                <div class="graph-title" :class="{'has-data':Object.keys(graph.data).length>0}">{{graph.title}}</div>
                <div class="graph-elapsed-time">
                    <div v-if="graph.elapsedTime">{{graph.elapsedTime}}s</div>
                </div>
            </div>
            <v-wait-modal :show="graph.inProgress" no-text class="wait-in-progress"/>
        </div>
    </div>
</template>

<script>
    import {mapGetters, mapMutations} from 'vuex';
    import {faCheck} from '@fortawesome/free-solid-svg-icons'
    import vWaitModal from '@/components/v-wait-modal.vue';

    export default {
        name: "v-toolbox-graph",
        data() {
            return {
                icons: {
                    select: faCheck,
                }
            }
        },
        components: {
            vWaitModal,
        },
        computed: {
            ...mapGetters(['graphs', 'selectedGraphName']),
        },
        methods: {
            ...mapMutations(['updateSelectedGraph'])
        }
    }
</script>

<style scoped lang="scss">
    $radius-size: 12px;
    .graph-selector-container {
        position: relative;
        padding: 1rem 0;
        height: calc(#{$toolbox-body-height} - #{$toolbox-body-border-bottom-width});
        display: flex;
        flex-flow: row wrap;
        justify-content: flex-start;
        align-content: flex-start;

        .graph-selector {
            flex: 0 0 calc(50% - 2rem);
            color: $dark-gray-color;
            border-radius: $radius-size;
            position: relative;
            margin: 1rem;
            height: 6rem;
            cursor: pointer;
            display: flex;
            flex-flow: row wrap;
            align-items: stretch;
            justify-content: stretch;
            box-shadow: 3px 3px 10px 2px darken($mid-gray-color, 20%);
            transition: transform 50ms linear;
            &.selected {
                box-shadow: 1px 1px 4px darken($mid-gray-color, 25%);
                transform: translate(3px, 3px);
            }
            .select-icon {
                flex: 0 0 auto;
                display: flex;
                align-items: center;
                justify-content: center;
                width: 15%;
                background-color: lighten($dark-gray-color, 18%);
                color: darken($white-color, 20%);
                border-radius: $radius-size 0 0 $radius-size;
                &.selected {
                    color: #adeaa6;
                    //background-color: #0caf0f;
                    background-color: #037b48;
                }
            }
            .graph-data {
                flex: 1 2 5rem;

                display: flex;
                flex-flow: column wrap;
                justify-items: flex-start;
                .graph-title,
                .graph-elapsed-time {
                    display: flex;
                    align-items: center;
                    justify-content: center;
                    justify-items: center;
                }
                .graph-title {
                    flex: 5 2 auto;
                    text-align: center;
                    border-radius: 0 $radius-size 0 0;
                    color: lighten($mid-gray-color, 10%);
                    font-size: .9rem;
                    padding: .3rem .5rem;
                    font-weight: bold;
                    background-color: lighten($dark-gray-color, 5%);
                    &.has-data {
                        background-color: #1e8e5e;
                        color: darken($white-color, 10%);
                    }
                }
                .graph-elapsed-time {
                    flex: 1 0 auto;
                    min-height: 1rem;
                    border-radius: 0 0 $radius-size 0;
                    font-size: .8rem;
                    font-weight: bold;
                    letter-spacing: .1rem;
                    color: $white-color;
                    background-color: lighten($dark-gray-color, 2%);
                }
            }
            .wait-in-progress {
                height: 100%;
                width: 100%;
                opacity: .6;
                border-radius: $radius-size;
            }
        }
    }

    @media all and (max-height: $sm__height-limit) {
        .graph-selector-container {
            height: calc(#{$toolbox-body-height__sm} - #{$toolbox-body-border-bottom-width});
        }
    }
</style>