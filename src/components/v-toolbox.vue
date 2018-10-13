<template>
    <div>
        <div :class="['toolbox',{'expanded':!!currentItem,'toggle-on':showToolbox}]">
            <div v-for="item in items"
                 :key="item.title"
                 @click="currentItem && currentItem.title===item.title ? currentItem=null: currentItem=item"
                 :class="['toolbox-item',{'active': currentItem && currentItem.title === item.title}]">
                <font-awesome-icon :icon="item.icon" size="2x" class="toolbox-item-icon"/>
                <div class="toolbox-item-title">{{item.title}}</div>
            </div>
            <div class="toolbox-item active" @click="toggleToolbox">
                <font-awesome-icon class="toolbox-item-icon" :icon="!showToolbox?arrowRightIcon:arrowLeftIcon"
                                   size="2x"/>
            </div>
        </div>

        <keep-alive>
            <div class="toolbox-body" v-if="!!currentItem">
                <transition name="component-switch" mode="out-in">
                    <component :is="!!currentItem ? currentItem.component : null"/>
                </transition>
            </div>
        </keep-alive>
    </div>
</template>

<script>
    import {
        faAngleDoubleLeft,
        faAngleDoubleRight,
        faChartBar,
        faExclamationTriangle,
        faFileCode,
        faFilter,
        faFlask,
    } from "@fortawesome/free-solid-svg-icons";
    import vToolboxFiles from "@/components/v-toolbox-files.vue";
    import vToolboxScrutinize from "@/components/v-toolbox-scrutinize.vue";
    import vToolboxFilter from "@/components/v-toolbox-filter.vue";
    import vToolboxGraph from "@/components/v-toolbox-graph.vue";
    import {mapGetters} from 'vuex';

    export default {
        name: "v-toolbox",
        data() {
            return {
                items: [
                    {title: "Files", icon: faFileCode, component: vToolboxFiles},
                    {title: "Scrutinize", icon: faFlask, component: vToolboxScrutinize},
                    {title: "Graphs", icon: faChartBar, component: vToolboxGraph},
                    {title: "Filters", icon: faFilter, component: vToolboxFilter},
                ],
                currentItem: null,
                showToolbox: false,
                arrowRightIcon: faAngleDoubleRight,
                arrowLeftIcon: faAngleDoubleLeft,
                warningIcon: faExclamationTriangle,
            }
        },
        computed: {
            ...mapGetters(['graphs']),
        },
        methods: {
            toggleToolbox() {
                this.showToolbox = !this.showToolbox;
                if (!this.showToolbox) {
                    this.currentItem = null;
                }
            }
        }
    }
</script>

<style lang="scss" scoped>
    .component-switch-enter-active, .component-switch-leave-active {
        transition: all .2s ease-out;
    }

    .component-switch-enter, .component-switch-leave-to {
        transform: translateX(-#{$toolbox-width * .7 });
        opacity: 0;
    }

    .toolbox {
        //$toolbox-left-indent: 20.5rem;
        $toolbox-left-indent: $toolbox-width * .83;
        position: fixed;
        top: $toolbox-top;
        left: -#{$toolbox-left-indent};
        height: $toolbox-height;
        width: $toolbox-width;
        background-color: $dark-gray-color;
        padding: .2rem .5rem;
        border-radius: 0 8px 8px 0;

        transition: transform .3s cubic-bezier(0.25, 0.58, 0.70, 0.96);

        display: flex;
        align-items: center;
        justify-content: space-around;
        z-index: 99;
        &.expanded {
            border-radius: 0 8px 0 0;
        }
        &.toggle-on {
            transform: translateX($toolbox-left-indent);
        }
        .toolbox-item {
            margin: .2rem;
            padding: .5rem;
            cursor: pointer;
            font-size: .8rem;

            display: flex;
            flex-direction: column;
            align-items: center;
            justify-content: center;

            color: lighten($mid-gray-color, 15%);
            background-color: transparent;

            transition: transform .15s ease;

            &:not(.active):hover {
                transform: translateY(-.15rem);
            }

            &.active {
                color: $white-color;
            }
            .toolbox-item-icon {
                height: 1.7rem;
                width: auto;
            }
            .toolbox-item-title {
                padding: .5rem .2rem 0;
                font-weight: bold;
            }
        }

    }

    .toolbox-body {
        position: fixed;
        left: 0;
        top: $toolbox-top + $toolbox-height;
        background-color: lighten($dark-gray-color, 9%);
        border: solid $toolbox-body-border-bottom-width $dark-gray-color;
        border-left: 0;
        border-top: 0;
        padding: 0;
        width: $toolbox-width;
        max-width: $toolbox-width;
        max-height: $toolbox-body-height;
        border-radius: 0 0 8px 0;
        font-size: .9rem;
        z-index: 99;
        min-height: $toolbox-body-height;
    }

    @media all and (max-height: $sm__height-limit) {
        .toolbox-body {
            max-height: $toolbox-body-height__sm;
            min-height: $toolbox-body-height__sm;
        }
    }
</style>
