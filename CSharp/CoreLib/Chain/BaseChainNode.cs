using System;

namespace CoreLib.Chain
{
    public abstract class BaseChainNode : IChainNode
    {
        public const string ERROR_KEY = "ChainError";

        public IChainNode m_NextNode;

        public void Invoke(SimChainContext ctx)
        {
            try
            {
                this.Node(ctx);

                if (this.m_NextNode != null)
                    this.m_NextNode.Invoke(ctx);
            }
            catch (Exception ex)
            {
                var errMsg = string.Format("{0}:{1}", this.GetType().Name, ex.Message);
                ctx.SetValue(ERROR_KEY, errMsg);
                Console.Write(errMsg);
            }
        }

        void IChainNode.Node(SimChainContext ctx)
        {
            this.Node(ctx);
        }

        protected abstract void Node(SimChainContext ctx);

        public IChainNode SetNext(IChainNode nextNode)
        {
            this.m_NextNode = nextNode;
            return this;
        }
    }
}
