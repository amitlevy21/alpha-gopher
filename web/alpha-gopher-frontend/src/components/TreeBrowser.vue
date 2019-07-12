<template>
  <div>
    <div 
      :style="{'margin-left': `${depth * 20}px`}"
      class="node"
      @click="nodeClicked"
    >
      <span 
        v-if="hasChildren"
        class="type"
      >{{ expanded ? '&#9660;' : '&#9658;' }}</span>
      <span
        v-else
        class="type"
      >&#9671;</span>
      <span>{{ node.name }}</span>
    </div>
    <div v-if="expanded">
      <TreeBrowser 
        v-for="child in node.contents"
        :key="child.name"
        :node="child"
        :depth="depth + 1"
        @onClick="(node) => $emit('onClick', node)"
      />
    </div>
  </div>
</template>

<script>
export default {
  name: 'TreeBrowser',
  props: {
    node: {
      type: Object,
      default: function() {}
    },
    depth: {
      type: Number,
      default: 0,
    }
  },
  data() {
    return {
      expanded: false,
    }
  },
  computed: {
    hasChildren() {
      return this.node.contents;
    }
  },
  methods: {
    nodeClicked() {
      this.expanded = !this.expanded;
      if (!this.hasChildren) {
        this.$emit('onClick', this.node);
      }
    }
  }
}
</script>

<style scoped>
.node {
  text-align: left;
  font-size: 18px;
}
.type {
  margin-right: 10px;
}
</style>