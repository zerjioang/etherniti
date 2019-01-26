<template>
    <generic
    :title="title"
    :message="message"
    :icon="icon"
    :buttonText="buttonText"
    v-on:close="$emit('close')">
    <div slot="content">
        <img class="minilogo" :src="require('@/assets/images/ethereum.png')">
        <img class="minilogo" :src="require('@/assets/images/quorum.png')">
        <p class="justified">
            <b>gaethway</b> allows companies, developers and startups to interact with Ethereum in a very easy and friendly way. We add an abstraction layer that allows a <b>fast</b>, reliable and <b>secure</b> interaction with Ethereum networks and their smart contracts. We provide a fully featured High Performance WebAPI to operate ethereum as its usually done with other tools such as <b>Web3js</b>, <b>Metamask</b>, etc.
        </p>
        <p>
            <b>gaethway</b> can be very useful in several scenarios, such as:
        </p>
        <ul>
            <li>
                Ethereum interaction learning from scratch. The concept of address, key, signature, etc.
            </li>
            <li>
                Add direct interaction from low resources embedded devices that cannot directly sign transactions, such as <b>ESP32</b>, <b>Arduino Uno</b>, etc.
            </li>
            <li>
                Add direct interaction from legacy software written in no compatible <b>Ethereum SDK</b>, leveraging the interaction in a <b>secure WebAPI</b>.
            </li>
            <li>
                Connect to the Blockchain some of the most used corporate <b>CRM, ERP, CMS</b>, etc.
            </li>
            <li>
                <b>Debug</b> software interaction, and check for errors before publishing your application or service.
            </li>
        </ul>
    </div>
    </generic>
</template>

<script>
export default {
    name: 'about-dialog',
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
        generic: () => import('@/components/dialog/generic')
    }
}
</script>

<style type="text/css" scoped="true">
p, ul, li .justified {
    font-size: 16px;
    padding-left: 30px;
    padding-right: 30px;
    padding-bottom: 30px;
    text-align: justify;
    color: #555;
}
ul {
    margin-left: 15px;
}
p {
    padding-bottom: 2px;
}
.minilogo {
    width: 80px;
    margin: 10px;
}
</style>