<template>
    <div ref="container" class="modal-container" v-if="show">
        <font-awesome-icon :icon="icons.waiting" size="3x" spin class="animated"/>
        <div class="waiting-message animated">Please Wait...</div>
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
    }
</script>

<style scoped lang="scss">
    .modal-container {
        color: lighten($dark-gray-color, 60%);
        position: absolute;
        height: calc(#{$toolbox-body-height} - #{$toolbox-body-border-bottom-width});
        width: calc(#{$toolbox-width} - #{$toolbox-body-border-bottom-width});
        opacity: .9;
        background-color: $dark-gray-color;
        z-index: 999;
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
            animation-name: fade-out-in;
            animation-iteration-count: infinite;
            animation-timing-function: ease-out;
            animation-duration: 2s;
        }
    }

    @media all and (max-height: $sm__height-limit) {
        .modal-container {
            height: calc(#{$toolbox-body-height__sm} - #{$toolbox-body-border-bottom-width});
        }
    }

    @keyframes fade-out-in {
        0%{
            opacity: .3;
        }
        50%{
            opacity: 1;
        }
        100%{
            opacity: .3;
        }
    }
</style>