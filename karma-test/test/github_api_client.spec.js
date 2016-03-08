import { Hoge } from '../src/test.js';

describe('test', () => {
  it('test', () => {
    expect(new Hoge().hoge()).toEqual('hoge');
  })
});
