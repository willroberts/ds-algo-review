// ArrayLists are dynamic arrays.
// There is no built-in synchronization, so concurrent access must be checked.
// The underlying array grows automatically by increasing by 1.5x at insertion time.
import java.util.ArrayList;

public class ArrayLists
{
	public static void main(String[] args)
	{
		ArrayList<Integer> a = new ArrayList<Integer>(0);
		a.add(1);
		a.add(2);
		a.add(3);
		a.add(4);
		a.add(0, 0);
		a.remove(4);
		for (int i: a) {
			System.out.println(i);
		}
	}
}
