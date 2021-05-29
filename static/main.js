import { h, render, Component } from "https://unpkg.com/preact@latest?module";
// In case you need hooks uncomment this line
// import {} from "https://unpkg.com/preact@latest/hooks/dist/hooks.module.js?module";
import htm from "https://unpkg.com/htm?module";

// Initialize htm with Preact
const html = htm.bind(h);

class WLTRow extends Component {
    constructor(props) {
        super(props);        
        this.state = { type: props.type, W: 0, L: 0, T: 0};
        console.log(props.type)
    }

    componentDidMount() {
        this.timer = setInterval(() => {
            fetch("/getResults?class=" + this.state.type)
                .then(resp => resp.json())
                .then((data) => {
                    this.setState({
                        W: data.wins,
                        L: data.losses,
                        T: data.ties
                    })
                })
        }, 1000);
    }

    componentWillUnmount() {
        clearInterval(this.timer);
    }

    render({type}) {
        return html`
        <tr class="result-row">
            <td class="result-score"><img class="image-role" src="static/${type}.png" /></td>
            <td class="result-score result-score-W">${this.state.W}</td>
            <td class="result-score">-</td>
            <td class="result-score result-score-T">${this.state.T}</td>
            <td class="result-score">-</td>
            <td class="result-score result-score-L">${this.state.L}</td>
        </tr>
        `
    }
}

render(
  html`
  <table class="result-table">
    <tr class="result-row">
      <th></th>
      <th class="result-header result-header-W">W</th>
      <th class="result-header">-</th>
      <th class="result-header result-header-T">T</th>
      <th class="result-header">-</th>
      <th class="result-header result-header-L">L</th>
    </tr>
    <${WLTRow} type="tank" />
    <${WLTRow} type="dps" />
    <${WLTRow} type="sup" />
  </table>
  `,
  document.querySelector("#root")
);
