class App extends React.Component {
  constructor(props) {
    super(props);
    this.state = { signIn: false };
  }
  changeFunc = () => {
    this.setState((state, props) => ({ signIn: !state.signIn }));
  };
  render() {
    return (
      <div className="containter">
        <div className="reg_form">
          <div className="auth_switcher">
            <label>
              <input
                defaultChecked="true"
                type="radio"
                name="auth_switcher"
                onChange={this.changeFunc}
              />
              <div>Зарегистрироваться</div>
            </label>
            <label>
              <input
                type="radio"
                name="auth_switcher"
                onChange={this.changeFunc}
              />
              <div>Войти</div>
            </label>
          </div>
          {this.state.signIn ? <Login /> : <Registration />}
        </div>
      </div>
    );
  }
}

class Registration extends React.Component {
  constructor(props) {
    super(props);
  }
  render() {
    return (
      <div>
        <h1 className="auth_header">Зарегистрируйтесь уже сейчас!</h1>
        <form className="auth_form" method="post" action="/registration">
          <input
            type="text"
            name="first_name"
            placeholder="Имя"
            autoComplete="off"
          />
          <input
            type="text"
            name="last_name"
            placeholder="Фамилия"
            autoComplete="off"
          />
          <input
            type="text"
            name="email"
            placeholder="Email адрес"
            autoComplete="off"
          />

          <select name="sex" className={"auth__sex_switcher"}>
            <option>Пол: Мужской</option>
            <option>Пол: Женский</option>
          </select>

          <input
            type="text"
            name="password"
            placeholder="Пароль"
            autoComplete="off"
            className="full_width"
          />
          <input
            type="submit"
            className="full_width"
            id="signup_submit_button"
            value="ЗАРЕГИСТРИРОВАТЬСЯ"
          />
        </form>
      </div>
    );
  }
}
class Login extends React.Component {
  constructor(props) {
    super(props);
  }
  render() {
    return (
      <div>
        <h1 className="auth_header">Приятно видеть вас снова!</h1>
        <form className="auth_form" method="post" action="/login">
          <input
            type="text"
            name="email"
            placeholder="Email адрес"
            autoComplete="off"
            className="full_width"
          />
          <input
            type="text"
            name="password"
            placeholder="Пароль"
            autoComplete="off"
            className="full_width"
          />
          <input
            type="submit"
            className="full_width"
            id="signup_submit_button"
            value="ВОЙТИ"
          />
        </form>
      </div>
    );
  }
}
ReactDOM.render( <App/>, document.querySelector("#root"));