class Header {
  public counter = 0;

  public increment(): number {
    return ++this.counter;
  }
}

export { Header };
