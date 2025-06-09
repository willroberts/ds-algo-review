// Vectors are essentially synchronized ArrayLists.
// The underlying array grows automatically by doubling at insertion time.
import java.util.Vector;

public class Vectors
{
	public static void main(String[] args)
	{
		Vector<Integer> v = new Vector<Integer>(0);
		v.addElement(1);
		v.addElement(2);
		v.addElement(3);
		v.addElement(4);
		v.insertElementAt(0, 0);
		v.removeElementAt(4);
		for (int i: v) {
			System.out.println(i);
		}
	}
}
