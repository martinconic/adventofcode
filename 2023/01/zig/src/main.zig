const std = @import("std");

pub fn main() !void {
    const allocator = std.heap.page_allocator;

    const file = try std.fs.cwd().openFile("../input/input.txt", .{});
    defer file.close();

    var buffered = std.io.bufferedReader(file.reader());
    var reader = buffered.reader();

    var arr = std.ArrayList(u8).init(allocator);
    defer arr.deinit();

    var sum: u32 = 0;

    while (true) {
        reader.streamUntilDelimiter(arr.writer(), '\n', null) catch |err| switch (err) {
            error.EndOfStream => break,
            else => return err,
        };

        const line = try arr.toOwnedSlice();
        defer allocator.free(line);

        var numbers = std.ArrayList(u32).init(std.heap.page_allocator);
        defer numbers.deinit();

        // Iterate over each character in the line
        for (line) |c| {
            // Check if the character is a digit
            if (std.ascii.isDigit(c)) {
                // Convert the character to an integer and add to the list
                _ = try numbers.append(@as(u32, @intCast(c - '0')));
            }
        }

        // Ensure there are enough numbers to perform the operation
        if (numbers.items.len >= 1) {
            sum += 10 * numbers.items[0] + numbers.items[numbers.items.len - 1];
        }
        arr.clearRetainingCapacity();
    }

    std.debug.print("Sum: {d}\n", .{sum});
}