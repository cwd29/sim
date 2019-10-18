namespace CoreLib.Chain
{
    public interface IChainNode
    {
        void Node(SimChainContext ctx);

        void Invoke(SimChainContext ctx);
    }
}
