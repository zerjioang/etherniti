<template>
    <div class="dialog">
        <div class="row">
            <div class="col-md-12 centered">
                <span class="material-icons top-big">{{icon}}</span>
                <h1 class="title centered">{{title}}</h1>
                <p class="message centered">{{message}}</p>
                <slot name="content"></slot>
                <button class="btn bg-indigo waves-effect btn-lg fit" @click="$emit('close')">{{buttonText}}</button>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: 'generic-dialog-message',
    props: {
        title: {
            type: String,
            default: "no-title"
        },
        message: {
            type: String,
            default: "no-message"
        },
        icon: {
            type: String,
            default: 'info_outline'
        },
        buttonText: {
            type: String,
            default: 'Close'
        }
    },
    data() {
        var data = {
        }
        return data;
    },
    methods: {
    },
    created(){

    },
    mounted(){

    },
    components: {

    }
}
</script>

<style type="text/css" scoped="true">

.title {
    font-size: 20px !important;
    margin: 0 0 5px 0 !important;
    color: #575757;
    text-align: center;
    font-weight: 600;
    text-transform: none;
    position: relative;
    margin: 25px 0;
    padding: 0;
    line-height: 40px;
    display: block;
}

.message {
    font-size: 16px;
    color: #555;
}

.top-big {
    font-size: 50pt;
    padding-top: 20px;
}

.centered {
    text-align: center;
}

.fit {
    font-size: 16px !important;
    margin-bottom: 16px;
}
</style>