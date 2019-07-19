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
      <span>
        {{ nodeName }}
      </span>
      <ul v-if="node.name != '/'">
        <li>
          User: {{ node.user }}
        </li>
        <li>
          Group: {{ node.group }}
        </li>
        <li>
          Last Changed: {{ node.time }}
        </li>
        
        <app-popup />
      </ul>
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
import Popup from "./Popup.vue";
export default {
  name: 'TreeBrowser',
  components: {
    'app-popup': Popup
  },
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
      dialog: false
    }
  },
  computed: {
    hasChildren() {
      return this.node.contents;
    },
    nodeName() {
      if (this.node.name == '/') {
        return '/'
      } else {
        let arr = this.node.name.split('/')
        return arr.pop()
      }
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
ul {
  list-style-type: none;
  margin: 4px;
  padding: 4px;
  overflow: hidden;
}
li {
  float: left;
  margin: 4px;
  padding: 0;
}
</style>