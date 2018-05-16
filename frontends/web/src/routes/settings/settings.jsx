import { Component } from 'preact';
import { translate } from 'react-i18next';
import { apiGet, apiPost } from '../../utils/request';
import { Button, Checkbox } from '../../components/forms';
import Toast from '../../components/toast/Toast';

@translate()
export default class Settings extends Component {
    state = {
        toast: false,
        config: null,
    }

    componentDidMount() {
        apiGet('config').then(config => this.setState({ config }));
    }

    toggleAccountActive = event => {
        let config = this.state.config;
        config.backend[event.target.id] = event.target.checked;
        this.setState({ config });
    }

    save = event => {
        if (!this.state.config) {
            return;
        }
        apiPost('config', this.state.config).then(() => {
            this.setState({ toast: true });
        });
    }

    render({
        t,
    }, {
        config, toast
    }) {
        return (
            <div class="container">
              <div class="innerContainer">
                <div class="header">
                  <h2>{t('settings.title')}</h2>
                </div>
                <div class="content flex flex-column flex-start">
                  { config && (
                      <div>
                        Active accounts:<br/>
                        <Checkbox
                          checked={config.backend.bitcoinP2PKHActive}
                          id="bitcoinP2PKHActive"
                          onChange={this.toggleAccountActive}
                          label="Bitcoin Legacy"
                          />
                        <Checkbox
                          checked={config.backend.bitcoinP2WPKHP2SHActive}
                          id="bitcoinP2WPKHP2SHActive"
                          onChange={this.toggleAccountActive}
                          label="Bitcoin Segwit"
                          />
                        <Checkbox
                          checked={config.backend.bitcoinP2WPKHActive}
                          id="bitcoinP2WPKHActive"
                          onChange={this.toggleAccountActive}
                          label="Bitcoin Native Segwit"
                          />
                        <Checkbox
                          checked={config.backend.litecoinP2WPKHP2SHActive}
                          id="litecoinP2WPKHP2SHActive"
                          onChange={this.toggleAccountActive}
                          label="Litecoin Segwit"
                          />
                      </div>
                  )
                  }
            </div>
                <Button primary onClick={this.save}>
                {t('button.save')}
            </Button>
                </div>
                <Toast
                    trigger={toast}
                    theme="success"
                    message="Settings saved. Please restart the application for the changes to take effect."
                    onHide={() => this.setState({ toast: false })}
                />
                </div>
        );
    }
}
