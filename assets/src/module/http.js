import axios from 'axios';

class Http {
  constructor() {
    this.client = axios.create({
      baseURL: 'http://localhost:8080',
      headers: {
        ContentType: 'application/json',
      },
      responseType: 'json',
    });
  }

  get(path) {
    return this.client({
      method: 'GET',
      url: path,
    }).catch(error => console.log(error));
  }

  post(path, data) {
    return this.client({
      method: 'POST',
      url: path,
      data,
    }).catch(error => console.log(error));
  }

  put(path, data) {
    return this.client({
      method: 'PUT',
      url: path,
      data,
    }).catch(error => console.log(error));
  }

  delete(path) {
    return this.client({
      method: 'DELETE',
      url: path,
    }).catch(error => console.log(error));
  }
}

export default new Http();
