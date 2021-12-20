const AttributeBinding = {
    data() {
        return {
            message: "You loaded this page on " + new Date().toLocaleTimeString()
        }
    }
}

Vue.createApp(AttributeBinding).mount("#bind-attribute")
