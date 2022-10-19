// Can I avoid this hardcoded path?
import "../../../../assets/node_modules/xterm/css/xterm.css"
import { Terminal } from "../../../../assets/node_modules/xterm"

export default IExTerminal = {
    mounted() {
        let term = new Terminal();
        term.open(this.el.querySelector(".xtermjs_container"));

        buffer = ""
        afterbuffer = ""

        term.write("> ")

        term.onKey((event) => {

            const BACKSPACE = 8;
            const ENTER = 13;
            const ARROW_LEFT = 37;
            const ARROW_RIGHT = 39;

            key = event.key;
            ev = event.domEvent;

            console.log(event);

            if (ev.keyCode == BACKSPACE) {
                buffer = buffer.slice(0, -1);
                term.write('\b \b');
                return;
            }

            if (ev.keyCode == ENTER) {
                console.log(buffer + afterbuffer);
                this.pushEventTo(this.el, "command", buffer + afterbuffer);
                buffer = "";
                afterbuffer = "";
                term.write("\r\n");
                return;
            }

            if (ev.keyCode == ARROW_LEFT) {
                afterbuffer = buffer.slice(-1) + afterbuffer;
                buffer = buffer.slice(0, -1);
                term.write(key);
                return;
            }

            if (ev.keyCode == ARROW_RIGHT) {
                buffer = buffer + afterbuffer.slice(0, 1);
                afterbuffer = afterbuffer.slice(1);
                term.write(key);
                return;
            }

            buffer += key
            term.write(key + afterbuffer);
            [...afterbuffer].forEach(c => term.write("\u001b[D")) // return the cursor to the right position
        })

        this.handleEvent("print", e => {
            console.log(e.data)
            term.write(e.data);
            term.write("\r\n> ");
        });
    }
}