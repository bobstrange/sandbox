import { Github } from 'github-api'

export class GithubClient {
  constructor() {
    const TOKEN = 'b56215b6b5b126ba2ab2fff56840828765ec865e';
    const AUTH  = 'oauth';
    const client = new Github(
      token: TOKEN,
      auth : AUTH
    )

    this._client = client;
  }

  fetchIssues(options) {
    new Promise((resolve, reject) => {
      this._issues().list(options)
    });
  }

  _issues() {
    this._client.getIssues(this._username, this.reponame);
  }
  // username can also be  a organization
  set username(username) {
    this._username = username;
  }
  set reponame(reponame) {
    this._reponame = reponame;
  }
}
