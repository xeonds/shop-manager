<template>
  <div class="customer-view">
    <el-row>
      <el-col :span="24">
        <h1>总览</h1>
      </el-col>
    </el-row>
    <el-row :gutter="20">
      <el-col :span="18">
        <el-alert title="开发中" type="warning" description="警告：创建订单功能正在开发中，请等待开发完成后再使用" show-icon class="card-row" />
        <el-card class="card-row">
          <template #header>
            <el-text>创建订单</el-text>
          </template>
          <el-form :model="product" ref="orderCreate" :label-width="120">
            <el-form-item label="商品名称">
              <el-select v-model="product.name" value-key="id" placeholder="请选择购买的商品">
                <el-option v-for="product in productList" :key="product.id" :label="product.name" :value="product" />
              </el-select>
            </el-form-item>
            <el-form-item label="品牌">
              <el-input v-model="product.brand" />
            </el-form-item>
            <el-form-item label="购买数量（斤）">
              <el-input-number v-model="product.quantity" :min="0" :precision="2" :step="0.10" />
            </el-form-item>
            <el-form-item label="单价">
              <el-input-number v-model="product.unit_price" :min="0" :precision="2" :step="0.10" />
            </el-form-item>
            <el-form-item label="总价">
              <el-input-number v-model="product.total_price" :min="0" :precision="2"
                :value="product.quantity * product.price" />
            </el-form-item>
            <el-form-item label="优惠金额">
              <el-input-number v-model="product.total_price" :min="0" :precision="2"
                :value="product.quantity * product.price" />
            </el-form-item>
            <el-form-item prop="paid" label="实收">
              <el-input-number v-model="product.paid" :min="0" :precision="2" :step="100" />
            </el-form-item>
            <el-form-item prop="payment" label="支付方式">
              <el-radio-group v-model="product.payment">
                <el-radio-button label="现金" value="现金" />
                <el-radio-button label="微信" value="微信" />
                <el-radio-button label="支付宝" value="支付宝" />
              </el-radio-group>
            </el-form-item>
            <el-form-item prop="note" label="备注">
              <el-input v-model="product.note" type="textarea" />
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="submitForm('orderCreate')" plain>保存</el-button>
            </el-form-item>
          </el-form>
        </el-card>
        <el-card class="card-row">
          <template #header>
            <el-text>客户充值</el-text>
          </template>
          <el-form :inline="true" :model="customer" ref="customerRecharge">
            <el-form-item label="客户姓名">
              <el-select v-model="customer" filterable remote reserve-keyword placeholder="输入以快速查找"
                :remote-method="remoteMethod" value-key="id" clearable>
                <el-option v-for="item  in  customerList" :key="item.id" :label="item.name" :value="item"
                  class="customer-selector">
                  <span class="left">姓名 {{ item.name }}</span>
                  <span class="right">手机号 {{ item.phone }}</span>
                </el-option>
              </el-select>
            </el-form-item>
            <el-form-item label="充值金额">
              <el-input-number v-model="customer.balance" :min="0" :precision="2" :step="100" />
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="submitForm('customerRecharge')" plain>保存</el-button>
            </el-form-item>
          </el-form>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card class="card-row">
          <template #header>
            <el-text>客流总览</el-text>
          </template>
          <el-row>
            <el-col :span="24">
              <el-statistic title="今日客流" :value="productList.length" />
            </el-col>
          </el-row>
        </el-card>
        <el-card class="card-row">
          <template #header>
            <el-text>订单总览</el-text>
          </template>
          <el-row>
            <el-col :span="12">
              <el-statistic title="今日订单数" :value="customerList.length" />
            </el-col>
            <el-col :span="12">
              <el-statistic title="今日营收额" :value="premierCustomers" />
            </el-col>
          </el-row>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import customerAPI from "../../api/customer";
import orderAPI from "../../api/order";
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
    men2Women() {
      return (
        this.customerList.filter((customer) => {
          customer.sex == "man";
        }).length +
        "/" +
        this.customerList.filter((customer) => {
          customer.sex == "man";
        }).length
      );
    },
    premierCustomers() {
      return this.customerList.filter((customer) => {
        customer.premier == true;
      }).length;
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
    async submitForm(formName) {
      try {
        await this.$refs[formName].validate()
        switch (formName) {
          case 'orderCreate':
            await orderAPI.createOrder(this.order)
            this.$message.success('订单创建成功')
            break
          case 'customerRecharge':
            this.customer.vip_card.balance += this.customer.balance
            await customerAPI.updateCustomer(this.customer.id, this.customer)
            this.$message.success('充值成功，余额为 ' + this.customer.vip_card.balance + ' 元')
            break
        }
      } catch (error) {
        this.$message.error('操作失败: ' + error || '未知错误')
      } finally {
        this.resetForm(formName)
      }
    },
    resetForm(formName) {
      this.$refs[formName].resetFields()
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

.customer-selector {
  .left {
    float: left;
  }

  .right {
    float: right;
    color: var(--el-text-color-secondary);
    font-size: 13px;
  }
}
</style>
