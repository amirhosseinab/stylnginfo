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
            <component :class="['toolbox-body',{'expanded':!!currentItem}]"
                       :is="!!currentItem ? currentItem.component : null"/>
        </keep-alive>
    </div>
</template>

<script>
    import {
        faAngleDoubleLeft,
        faAngleDoubleRight,
        faChartBar,
        faFileCode,
        faFilter,
        faFlask
    } from "@fortawesome/free-solid-svg-icons";
    import vToolboxFiles from "@/components/v-toolbox-files.vue";
    import vToolboxAnalyze from "@/components/v-toolbox-analyze.vue";
    import vToolboxFilter from "@/components/v-toolbox-filter.vue";
    import vToolboxGraph from "@/components/v-toolbox-graph.vue";

    export default {
        name: "v-toolbox",
        data() {
            return {
                items: [
                    {title: "Files", icon: faFileCode, component: vToolboxFiles},
                    {title: "Analyze", icon: faFlask, component: vToolboxAnalyze},
                    {title: "Graphs", icon: faChartBar, component: vToolboxGraph},
                    {title: "Filters", icon: faFilter, component: vToolboxFilter},
                ],
                currentItem: null,
                showToolbox: false,
                arrowRightIcon: faAngleDoubleRight,
                arrowLeftIcon: faAngleDoubleLeft,
            }
        },

        methods: {
            toggleToolbox() {
                this.showToolbox = !this.showToolbox;
                if (!this.showToolbox) {
                    this.currentItem = null;
                }
            },
        }
    }
</script>

<style lang="scss" scoped>
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
        &.expanded {
            min-height: $toolbox-body-height;
        }
    }

    @media all and (max-height: $sm__height-limit) {
        .toolbox-body {
            max-height: $toolbox-body-height__sm;
            &.expanded {
                min-height: $toolbox-body-height__sm;
            }
        }
    }
</style>
