using CoreLib.Chain;
using NUnit.Framework;

namespace CoreLib.Tests.Chain
{
    class MockChainNode1 : BaseChainNode
    {
        public const string ERROR = "Error Test";

        public const string HAS_ERROR = "HasError";

        public const string ID = "MockChainNode1";

        public const string Key = "MOCK_CHAIN_NODE_1";

        protected override void Node(SimChainContext ctx)
        {
            if (ctx.HasKey(HAS_ERROR))
                throw new System.Exception(ERROR);

            ctx.SetValue(Key, ID);
        }
    }

    class MockChainNode2 : BaseChainNode
    {

        public const string Key = "MOCK_CHAIN_NODE_2";

        public const string ID = "MockChainNode2";

        protected override void Node(SimChainContext ctx)
        {
            ctx.SetValue(Key, ID);
        }
    }


    [TestFixture]
    class BaseChainNodeTests
    {
        [Test]
        public void OneNodeTest()
        {
            var ctx = new SimChainContext();
            new MockChainNode1().Invoke(ctx);

            Assert.IsTrue(ctx.HasKey(MockChainNode1.Key));
            Assert.AreEqual(ctx[MockChainNode1.Key], MockChainNode1.ID);
        }

        [Test]
        public void ChainTest()
        {
            var ctx = new SimChainContext();

            new MockChainNode1().SetNext(
                new MockChainNode2()
            ).Invoke(ctx);

            Assert.IsTrue(ctx.HasKey(MockChainNode1.Key));
            Assert.AreEqual(MockChainNode1.ID, ctx[MockChainNode1.Key]);

            Assert.IsTrue(ctx.HasKey(MockChainNode2.Key));
            Assert.AreEqual(MockChainNode2.ID, ctx[MockChainNode2.Key]);
        }

        [Test]
        public void ChainHasError()
        {
            var ctx = new SimChainContext();
            ctx.SetValue(MockChainNode1.HAS_ERROR, "");

            new MockChainNode1().SetNext(
                new MockChainNode2()
            ).Invoke(ctx);

            Assert.IsFalse(ctx.HasKey(MockChainNode1.Key));
            Assert.IsFalse(ctx.HasKey(MockChainNode2.Key));

            Assert.IsTrue(ctx.HasKey(BaseChainNode.ERROR_KEY));
            var errMsg = string.Format("MockChainNode1:{0}", MockChainNode1.ERROR);
            Assert.AreEqual(errMsg, ctx[BaseChainNode.ERROR_KEY]);
        }
    }
}
