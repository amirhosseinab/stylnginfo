<template>
    <div ref="container" class="modal-container" v-if="show">
        <font-awesome-icon :icon="icons.waiting" size="3x" spin/>
        <div class="waiting-message">Please Wait...</div>
    </div>
</template>

<script>
    import {faSpinner} from '@fortawesome/free-solid-svg-icons';

    export default {
        name: "v-wait-modal",
        data() {
            return {
                icons: {
                    waiting: faSpinner,
                }
            }
        },
        props: {
            show: {
                type: Boolean,
                required: true,
            }
        },
        updated() {
            let that = this;
            setTimeout(function () {
                that.$refs.container && that.$refs.container.classList.add("show")
            }, 10)
        }
    }
</script>

<style scoped lang="scss">
    .modal-container {
        color: $light-blue-color;
        position: absolute;
        height: calc(#{$toolbox-body-height} - #{$toolbox-body-border-bottom-width});
        width: $toolbox-width;
        opacity: .1;
        background-color: $dark-gray-color;
        z-index: 110;
        top: 0;
        left: 0;

        display: flex;
        flex-flow: row wrap;
        align-items: center;
        justify-content: center;

        transition: all .5s ease-out;
        .waiting-message {
            margin: 1rem;
            font-size: 1.5rem;
            font-weight: bold;
        }
        &.show {
            opacity: .9;
        }
    }

    @media all and (max-height: $sm__height-limit) {
        .modal-container {
            height: calc(#{$toolbox-body-height__sm} - #{$toolbox-body-border-bottom-width});
        }
    }
</style>