interface LinkedNode<K> {
  id: K
  prev: K
  next: K
}

export const linkedListMapToArray = <K>(
  map: Map<K, LinkedNode<K>>
): LinkedNode<K>[] => {
  const list: LinkedNode<K>[] = []
  const first = [...map.values()].find(p => p.prev === null)

  let now = first ?? null
  while (now !== null) {
    list.push(now)
    if (now.next) {
      now = map.get(now.next) ?? null
    } else {
      now = null
    }
  }

  return list
}
