const mutations = {
    setToken(state, token) {
      state.token = token;
    },
    setLogin(state, value) {
      state.loggedIn = value;
    },
    setLoginLink(state, value) {
      state.logLink = value;
    },
    reset(state) {
      state.loggedIn = false;
      state.token = "";
    },
  };
  
  export { mutations };