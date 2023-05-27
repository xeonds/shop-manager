<template>
  <div class="customer-view">
    <el-row>
      <el-col :span="24">
        <h1>财务</h1>
        <el-alert title="开发中" type="warning" description="警告：功能正在开发中，请等待开发完成后再使用" show-icon />
      </el-col>
    </el-row>
    <el-row :gutter="20">
      <el-col :span="18">
        <el-card class="card-row">
          <template #header>
            <el-text>营收额统计</el-text>
          </template>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card class="card-row">
          <template #header>
            <el-text>商品总览</el-text>
          </template>
          <el-row>
            <el-col :span="24">
              <el-statistic title="商品总量" :value="productList.length" />
            </el-col>
          </el-row>
        </el-card>
        <el-card class="card-row">
          <template #header>
            <el-text>顾客总览</el-text>
          </template>
          <el-row>
            <el-col :span="12">
              <el-statistic title="所有顾客" :value="customerList.length" />
            </el-col>
            <el-col :span="12">
              <el-statistic title="会员顾客" :value="premierCustomers" />
            </el-col>
          </el-row>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import customerAPI from "../../api/customer";
import productAPI from "../../api/product";

export default {
  data() {
    return {
      customer: {},
      product: {},
      customerList: [],
      productList: [],
    };
  },
  computed: {
    filteredCustomerList() {
      return this.customerList.filter((customer) => {
        return (
          customer.name.toLowerCase().includes(this.search.toLowerCase()) ||
          customer.email.toLowerCase().includes(this.search.toLowerCase()) ||
          customer.phone.toLowerCase().includes(this.search.toLowerCase()) ||
          customer.id.toString().includes(this.search)
        );
      });
    },
    premierCustomers() {
      return this.customerList.filter(customer => customer.is_vip == true).length
    },
  },
  methods: {
    async showCustomerList() {
      await customerAPI.getCustomers().then((res) => {
        this.customerList = res;
      }).catch((err) => {
        this.$message.error("顾客数据加载失败");
      });
    },
    async showProdctList() {
      await productAPI.getProducts().then((res) => {
        this.productList = res;
      }).catch((err) => {
        this.$message.error("商品数据加载失败");
      });
    },
    querySearch(queryString, cb) {
      const results = queryString
        ? this.productList.filter((product) =>
          product.name.toLowerCase().includes(queryString.toLowerCase())
        )
        : this.productList;
      cb(results);
    },
    handleSelect(item) {
      this.product = item;
    },
  },
  mounted() {
    this.showCustomerList();
    this.showProdctList();
  },
};
</script>

<style lang="less" scoped>
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.card-row {
  margin-bottom: 2rem;
}

.customer-view {
  margin: 0 auto;

  .el-row {
    margin-bottom: 2rem;
  }
}
</style>
