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
  };
  
  export { mutations };