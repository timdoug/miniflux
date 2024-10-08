class DomHelper {
    static isVisible(element) {
        return element.offsetParent !== null;
    }

    static openNewTab(url) {
        const win = window.open("");
        win.opener = null;
        win.location = url;
        win.focus();
    }

    static scrollPageTo(element, evenIfOnScreen) {
        window.scrollTo(0, element.offsetTop - 10);
    }

    static getVisibleElements(selector) {
        const elements = document.querySelectorAll(selector);
        return [...elements].filter((element) => this.isVisible(element));
    }

    static hasPassiveEventListenerOption() {
        var passiveSupported = false;

        try {
            var options = Object.defineProperty({}, "passive", {
                get: function() {
                    passiveSupported = true;
                }
            });

            window.addEventListener("test", options, options);
            window.removeEventListener("test", options, options);
        } catch(err) {
            passiveSupported = false;
        }

        return passiveSupported;
    }
}
